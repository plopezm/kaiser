package config

import (
	"log"
	"os"

	"github.com/plopezm/kaiser/utils"
)

// Configuration The configuration object
var Configuration ConfigurationData

func init() {
	err := utils.GetJSONObjectFromFile("kaiser.config.json", &Configuration)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	createWorkspace()
}

func createWorkspace() {
	log.Println("[config.go] Creating workspace if it does not exist in", Configuration.Workspace)
	os.Mkdir(Configuration.Workspace, 777)
}
