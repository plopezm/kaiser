package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core/types"
	"github.com/stretchr/testify/assert"
)

func Test_startNotifier_No_Workspace_Found(t *testing.T) {
	// Given
	config.Configuration = config.ConfigurationData{
		Workspace: "",
		LogFolder: "logs",
	}

	err := startNotifier()

	// Then
	assert := assert.New(t)
	assert.NotNil(err)
}

func Test_parseFolder_Workspace_Not_Found(t *testing.T) {
	// Given
	config.Configuration = config.ConfigurationData{
		Workspace: "",
		LogFolder: "./",
	}

	err := parseFolder(config.Configuration.Workspace)

	// Then
	assert := assert.New(t)
	assert.NotNil(err)
}

func setup() {
	os.Mkdir("workspace", 0700)
	var job = new(types.Job)
	types.InitializeJob(job)
	script := ""
	scriptFile := "testing.job.json"
	job.Name = "Testing"
	job.Tasks = make(map[string]*types.JobTask)
	job.Tasks["testTask"] = &types.JobTask{
		Name:       "testTask",
		OnFailure:  "",
		OnSuccess:  "",
		Script:     &script,
		ScriptFile: &scriptFile,
	}
	bytes, err := json.Marshal(job)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = ioutil.WriteFile("workspace/testing.job.json", bytes, 0644)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func afterTest() {
	os.RemoveAll("workspace")
}

func Test_parseFolder_Workspace(t *testing.T) {
	// Given
	setup()
	config.Configuration = config.ConfigurationData{
		Workspace: "./",
		LogFolder: "./",
	}

	err := parseFolder(config.Configuration.Workspace)

	// Then
	assert := assert.New(t)
	assert.Nil(err)
	afterTest()
}
