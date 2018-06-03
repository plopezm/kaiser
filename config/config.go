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
	configureLogger()
	createWorkspace()
}

func configureLogger() {
	f, err := os.OpenFile("logs/kaiser.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Mkdir("logs", 777)
		configureLogger()
	}
	log.SetOutput(f)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.LUTC)
	log.Println("========= Starting Kaiser =========")
}

func createWorkspace() {
	log.Println("Creating workspace if it does not exist in", Configuration.Workspace)
	os.Mkdir(Configuration.Workspace, 777)
}
