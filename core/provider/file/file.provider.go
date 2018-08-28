package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core/types"
	"github.com/plopezm/kaiser/core/validation"
	"github.com/plopezm/kaiser/utils"
)

// Channel the channel used to notify new jobs
var Channel chan types.Job

func init() {
	// This should prepare everything for thread looking for new files
	Channel = make(chan types.Job)
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
			err := parseJob(folderName+"/", f.Name())
			if err != nil {
				log.Println("Error parsing job file ["+f.Name()+"]:", err)
			}
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
func parseJob(folder string, filename string) error {
	var newJob types.Job
	hash, err := utils.GetJSONObjectFromFileWithHash(folder+filename, &newJob)
	if err != nil {
		return err
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
	err = validation.VerifyJob(&newJob)
	if err != nil {
		return err
	}
	Channel <- newJob
	return nil
}
