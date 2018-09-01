package file

import (
	"testing"

	"github.com/plopezm/kaiser/config"
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
