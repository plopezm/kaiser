package main

import (
	"log"

	_ "github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
	_ "github.com/plopezm/kaiser/interfaces/graphql"
	_ "github.com/plopezm/kaiser/interfaces/rest"
	_ "github.com/plopezm/kaiser/plugins/http"
	_ "github.com/plopezm/kaiser/plugins/logger"
	_ "github.com/plopezm/kaiser/plugins/system"
)

func main() {
	log.Println("========= Starting Kaiser =========")
	go core.StartServer(8080)
	engineInstance := core.GetEngineInstance()
	engineInstance.Start()
}
