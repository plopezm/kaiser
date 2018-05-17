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
		// log.Printf("[FileParser] Checking files in workspace \"%s\"\n", config.Configuration.Workspace)
		files, err := ioutil.ReadDir(config.Configuration.Workspace)
		if err != nil {
			log.Fatal(err)
		}

		// log.Println("Files found: ")
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), "job.json") {
				// log.Println(f.Name())
				parseJob(config.Configuration.Workspace + "/" + f.Name())
			}
		}
		time.Sleep(5000 * time.Millisecond)
	}
}

// parseJob Parses and creates a new job
func parseJob(filename string) {
	var newJob engine.Job
	err := utils.GetJSONObjectFromFile(filename, &newJob)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(parser.Observers) == 0 {
		return
	}
	_, ok := parser.jobs[newJob.Name]
	if !ok {
		parser.jobs[newJob.Name] = newJob
		parser.Notify(newJob)
	}
}
