package main

import (
	_ "github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
	_ "github.com/plopezm/kaiser/interfaces/graphql"
	_ "github.com/plopezm/kaiser/interfaces/rest"
)

func main() {
	go core.StartServer(8080)
	engineInstance := core.GetEngineInstance()
	engineInstance.Start()
}
