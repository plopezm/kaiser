package file

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/plopezm/kaiser/utils"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core/engine"
	"github.com/plopezm/kaiser/utils/observer"
)

// JobParser Is a parser who gets the jobs from workspace
type JobParser struct {
	observer.MapPublisher
	jobs map[string]engine.Job
}

// GetJobs Returns all registered jobs
func (parser *JobParser) GetJobs() map[string]engine.Job {
	return parser.jobs
}

var parser *JobParser

func init() {
	// This should prepare everything for thread looking for new files
	parser = new(JobParser)
	parser.jobs = make(map[string]engine.Job)
	parser.Observers = make(map[observer.Observer]struct{})
	go startParserScan()
}

// GetParser Returns the an instance of a FileJobParser
func GetParser() *JobParser {
	return parser
}

// StartParserScan Starts folder scan
func startParserScan() {
	for {
		parseFolder(config.Configuration.Workspace)
		time.Sleep(5000 * time.Millisecond)
	}
}

func parseFolder(folderName string) {
	time.Sleep(1000 * time.Millisecond)
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("checking folder: ", folderName)
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), "job.json") {
			parseJob(folderName+"/", f.Name())
		} else if f.IsDir() && isNotKaiserDir(f.Name()) {
			parseFolder(folderName + "/" + f.Name())
		}
	}
}

func isNotKaiserDir(folderName string) bool {
	return folderName != "disabled" && folderName != "plugins"
}

// parseJob Parses and creates a new job
func parseJob(folder string, filename string) {
	var newJob engine.Job
	err := utils.GetJSONObjectFromFile(folder+filename, &newJob)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(parser.Observers) == 0 {
		return
	}
	_, ok := parser.jobs[newJob.Name]
	if !ok {
		newJob.Folder = folder
		parser.jobs[newJob.Name] = newJob
		parser.Notify(newJob)
	}
}
