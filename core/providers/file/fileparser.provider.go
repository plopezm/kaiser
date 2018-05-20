package file

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/plopezm/kaiser/utils"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core/engine"
)

var parser *JobParser

func init() {
	// This should prepare everything for thread looking for new files
	parser = new(JobParser)
	parser.jobs = make(map[string]engine.Job)
	parser.Channel = make(chan engine.Job)
	go startProvider()
}

// JobParser Is a parser who gets the jobs from workspace
type JobParser struct {
	Channel chan engine.Job
	jobs    map[string]engine.Job
}

// GetJobs Returns all registered jobs
func (parser *JobParser) GetJobs() map[string]engine.Job {
	return parser.jobs
}

// GetParser Returns the an instance of a FileJobParser
func GetParser() *JobParser {
	return parser
}

// StartParserScan Starts folder scan
func startProvider() {
	for {
		parseFolder(config.Configuration.Workspace)
		time.Sleep(2000 * time.Millisecond)
	}
}

// parseFolder Scans the folders in the workspace
func parseFolder(folderName string) {
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

// isNotKaiserDir Returns true or false depending on the folder found (this is usefull for Kaiser reserved folders)
func isNotKaiserDir(folderName string) bool {
	return folderName != "disabled" && folderName != "plugins"
}

// parseJob Parses and creates a new job file
func parseJob(folder string, filename string) {
	var newJob engine.Job
	err := utils.GetJSONObjectFromFile(folder+filename, &newJob)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, ok := parser.jobs[newJob.Name]
	if !ok {
		newJob.Folder = folder
		parser.jobs[newJob.Name] = newJob
		parser.Channel <- newJob
	}
}
