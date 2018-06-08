package main

import (
	_ "github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/interfaces/graphql"
)

func main() {
	engine := core.New()
	go graphql.StartServer(8080)
	engine.Start()
}
