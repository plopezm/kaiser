package config

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConfigurationSuite(t *testing.T) {
	t.Run("Initialize config with a given configuration", whenInitializeConfigCreatesConfigFromAGivenConfiguration)
	t.Run("Creating workspace from a empty configuration", whenCreateWorkspaceEmptyConfiguration)
	t.Run("Create Workspace With Workspace Set In Config", whenCreateWorkspaceWithWorkspaceSetInConfig)
	t.Run("Configure Logger with defined log folder", whenConfigureLoggerWithDefinedLogFolder)
	t.Run("Configure Logger with undefined log folder", whenConfigureLoggerWithUndefinedLogFolder)
}

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

func whenInitializeConfigCreatesConfigFromAGivenConfiguration(t *testing.T) {
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

func whenCreateWorkspaceEmptyConfiguration(t *testing.T) {
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

func whenCreateWorkspaceWithWorkspaceSetInConfig(t *testing.T) {
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

func whenConfigureLoggerWithDefinedLogFolder(t *testing.T) {
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

func whenConfigureLoggerWithUndefinedLogFolder(t *testing.T) {
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
