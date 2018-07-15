package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/utils"
)

var Channel chan core.Job

func init() {
	// This should prepare everything for thread looking for new files
	Channel = make(chan core.Job)
	go startNotifier()
}

// StartParserScan Starts folder scan
func startNotifier() {
	for {
		parseFolder(config.Configuration.Workspace)
		time.Sleep(5000 * time.Millisecond)
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
	var newJob core.Job
	hash, err := utils.GetJSONObjectFromFileWithHash(folder+filename, &newJob)
	if err != nil {
		return
	}
	newJob.Folder = folder
	newJob.Hash = hash
	for key, task := range newJob.Tasks {
		task.Name = key
		if task.ScriptFile != nil {
			raw, err := ioutil.ReadFile(folder + *task.ScriptFile)
			if err != nil {
				log.Fatalln(err.Error())
				os.Exit(1)
			}
			fileContent := string(raw)
			task.Script = &fileContent
		}
	}
	Channel <- newJob
}
