package main

import (
	"log"

	"github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
)

func main() {
	log.Println(config.Configuration)
	engine := core.New()
	engine.Start()
}
