package main

import (
	_ "github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
	_ "github.com/plopezm/kaiser/interfaces/graphql"
)

func main() {
	go core.StartServer(8080)
	engineInstance := core.New()
	engineInstance.Start()
}
