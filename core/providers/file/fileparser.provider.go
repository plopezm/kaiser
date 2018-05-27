package file

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/plopezm/kaiser/utils"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core/engine"
)

var provider *JobProvider

func init() {
	// This should prepare everything for thread looking for new files
	provider = new(JobProvider)
	provider.jobs = make(map[string]*JobMetadata)
	provider.Channel = make(chan engine.Job)
	go startProvider()
}

// GetProvider Returns the an instance of a FileJobProvider
func GetProvider() *JobProvider {
	return provider
}

// JobProvider Is a parser who gets the jobs from workspace
type JobProvider struct {
	Channel chan engine.Job
	jobs    map[string]*JobMetadata
}

// GetJobs Returns all current jobs
func (prov *JobProvider) GetJobs() map[string]*JobMetadata {
	return prov.jobs
}

// JobMetadata represents a Job
type JobMetadata struct {
	hash   []byte
	ticker *time.Ticker
	job    *engine.Job
	Done   chan bool
}

// StartParserScan Starts folder scan
func startProvider() {
	for {
		parseFolder(config.Configuration.Workspace)
		time.Sleep(1000 * time.Millisecond)
	}
}

// parseFolder Scans the folders in the workspace
func parseFolder(folderName string) {
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), "job.json") {
			parseJob(folderName+"/", f.Name())
		} else if f.IsDir() && isNotKaiserDir(f.Name()) {
			parseFolder(folderName + "/" + f.Name())
		}
	}
}

// isNotKaiserDir Returns true or false depending on the folder found (this is usefull for Kaiser reserved folders)
func isNotKaiserDir(folderName string) bool {
	return folderName != "disabled" && folderName != "plugins"
}

// parseJob Parses and creates a new job file
func parseJob(folder string, filename string) {
	var newJob engine.Job
	hash, err := utils.GetJSONObjectFromFileWithHash(folder+filename, &newJob)
	if err != nil {
		return
	}

	storedJobMetadata, ok := provider.jobs[newJob.Name]

	// If the job has changed, the we should check it
	if !ok || bytes.Compare(storedJobMetadata.hash, hash) != 0 {
		if ok {
			storedJobMetadata.Done <- true
		}

		newJob.Folder = folder

		newJobMetadata := &JobMetadata{
			hash: hash,
			job:  &newJob,
			Done: make(chan bool),
		}

		checkPeriodicity(newJobMetadata)
		provider.jobs[newJobMetadata.job.Name] = newJobMetadata
		provider.Channel <- *newJobMetadata.job
	}
}

func checkPeriodicity(newJobMetadata *JobMetadata) {
	if len(newJobMetadata.job.Duration) > 0 {
		duration := utils.ParseDuration(newJobMetadata.job.Duration)
		if duration > 0 {
			newJobMetadata.ticker = time.NewTicker(duration)
			go periodHandler(newJobMetadata)
		}
	}
}

func periodHandler(jobData *JobMetadata) {
	for {
		select {
		case <-jobData.ticker.C:
			provider.Channel <- *jobData.job
		case <-jobData.Done:
			jobData.ticker.Stop()
			return
		}
	}
}
