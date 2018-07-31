package main

import (
	_ "github.com/plopezm/kaiser/config"
	core "github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/interfaces/graphql"
)

func main() {
	engineInstance := core.New()
	go graphql.StartServer(8080)
	engineInstance.Start()
}
