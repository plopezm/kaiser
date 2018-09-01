package config

import (
	"log"
	"os"
	"sync"

	"github.com/plopezm/kaiser/utils"
)

// Configuration The configuration object
var Configuration ConfigurationData
var once sync.Once

// InitializeConfig Initializes the configuration
func InitializeConfig(configFile string) {
	once.Do(func() {
		err := utils.GetJSONObjectFromFile(configFile, &Configuration)
		if err != nil {
			Configuration = ConfigurationData{
				Workspace: "workspace",
				LogFolder: "logs",
			}
		}
		err = configureLogger()
		if err != nil {
			log.Println(err)
		}
		createWorkspace()
	})
}

func configureLogger() error {
	if Configuration.LogFolder == "" {
		Configuration.LogFolder = "logs"
	}
	os.Mkdir(Configuration.LogFolder, 0700)
	f, err := os.OpenFile(Configuration.LogFolder+"/kaiser.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.LUTC)
	return nil
}

func createWorkspace() error {
	if Configuration.Workspace == "" {
		Configuration.Workspace = "workspace"
	}
	os.Mkdir(Configuration.Workspace, 0700)
	return nil
}
