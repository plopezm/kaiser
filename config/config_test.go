package config

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func afterTest() {
	os.RemoveAll("logs")
	os.RemoveAll("workspace")
}

func containsFileName(s []os.FileInfo, e string) bool {
	for _, a := range s {
		if a.Name() == e {
			return true
		}
	}
	return false
}

func Test_whenInitializeConfigCreatesConfigFromAGivenConfiguration(t *testing.T) {
	// Given
	Configuration = ConfigurationData{
		Workspace: "workspace",
		LogFolder: "logs",
	}

	// When
	InitializeConfig("testing")

	// Then
	assert := assert.New(t)
	currentDirFiles, err := ioutil.ReadDir(".")
	assert.Nil(err)

	assert.Equal(true, containsFileName(currentDirFiles, "logs"))
	assert.Equal(true, containsFileName(currentDirFiles, "workspace"))
	afterTest()
}

func Test_whenCreateWorkspaceEmptyConfiguration(t *testing.T) {
	// Given
	Configuration = ConfigurationData{
		Workspace: "",
		LogFolder: "",
	}

	// When
	err := createWorkspace()

	// Then
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Configuration.Workspace, "workspace")
	afterTest()
}

func Test_whenCreateWorkspaceWithWorkspaceSetInConfig(t *testing.T) {
	// Given
	Configuration = ConfigurationData{
		Workspace: "workspace",
	}

	// When
	err := createWorkspace()

	// Then
	assert := assert.New(t)
	assert.Nil(err)
	afterTest()
}

func Test_whenConfigureLoggerWithDefinedLogFolder(t *testing.T) {
	// Given
	Configuration = ConfigurationData{
		LogFolder: "logs",
	}

	// When
	err := configureLogger()

	// Then
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(log.Lshortfile|log.Ldate|log.Ltime|log.LUTC, log.Flags())
	afterTest()
}

func Test_whenConfigureLoggerWithUndefinedLogFolder(t *testing.T) {
	// Given
	Configuration = ConfigurationData{
		LogFolder: "",
	}

	// When
	err := configureLogger()

	// Then
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(Configuration.LogFolder, "logs")
	afterTest()
}
