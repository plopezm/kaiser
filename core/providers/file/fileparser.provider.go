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

var parser *JobProvider

func init() {
	// This should prepare everything for thread looking for new files
	parser = new(JobProvider)
	parser.jobs = make(map[string]*JobData)
	parser.Channel = make(chan engine.Job)
	go startProvider()
}

// GetParser Returns the an instance of a FileJobParser
func GetParser() *JobProvider {
	return parser
}

// JobProvider Is a parser who gets the jobs from workspace
type JobProvider struct {
	Channel chan engine.Job
	jobs    map[string]*JobData
}

// JobData represents a Job
type JobData struct {
	hash   []byte
	ticker *time.Ticker
	job    *engine.Job
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

	storedJob, ok := parser.jobs[newJob.Name]

	// If the job has changed, the we should check it
	if !ok || bytes.Compare(storedJob.hash, hash) != 0 {
		newJob.Folder = folder

		newJobData := &JobData{
			hash: hash,
			job:  &newJob,
		}

		if len(newJob.Duration) > 0 {
			duration := utils.ParseDuration(newJob.Duration)
			if duration > 0 {
				newJobData.ticker = time.NewTicker(duration)
				go func() {
					for range newJobData.ticker.C {
						parser.Channel <- *newJobData.job
					}
				}()
			}
		}

		parser.jobs[newJob.Name] = newJobData
		parser.Channel <- newJob
	}
}
