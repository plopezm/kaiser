package main

import (
	_ "github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/interfaces/cli"
)

func main() {
	engine := core.New()
	go cli.StartUICli()
	engine.Start()
}
