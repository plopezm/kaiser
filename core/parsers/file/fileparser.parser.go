package file

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core/engine"
	"github.com/plopezm/kaiser/core/parsers"
	"github.com/plopezm/kaiser/utils/observer"
)

// FileJobParser Is a parser who gets the jobs from workspace
type FileJobParser struct {
	observer.MapPublisher
	jobs map[string]engine.Job
}

// GetJobs Returns all registered jobs
func (parser *FileJobParser) GetJobs() map[string]engine.Job {
	return parser.jobs
}

// FileJobParserEvent Represents a FileJobParser event used in observers
type FileJobParserEvent struct {
	jobsFound []engine.Job
}

// Code Returns the event code. EventCodes can be checked in core/parsers/parsers.codes.go
func (e FileJobParserEvent) Code() int {
	return parsers.EventCodeFileParser
}

func init() {
	// This should prepare everything for thread looking for new files
	parser = new(FileJobParser)
	parser.jobs = make(map[string]engine.Job)
	parser.Observers = make(map[observer.Observer]struct{})
	go startParserScan()
}

var parser *FileJobParser

// GetParser Returns the an instance of a FileJobParser
func GetParser() *FileJobParser {
	return parser
}

func startParserScan() {
	for {
		log.Printf("[FileParser] Checking files in workspace \"%s\"\n", config.Configuration.Workspace)
		files, err := ioutil.ReadDir(config.Configuration.Workspace)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Files found: ")
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), "job.json") {
				log.Println(f.Name())
				parseJob(f.Name())
			}
		}
		time.Sleep(5000 * time.Millisecond)
	}
}

// parseJob Parses and creates a new job
func parseJob(filename string) {

}
