package main

import (
	_ "github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core/engine"
	"github.com/plopezm/kaiser/interfaces/graphql"
)

func main() {
	engineInstance := engine.New()
	go graphql.StartServer(8080)
	engineInstance.Start()
}
