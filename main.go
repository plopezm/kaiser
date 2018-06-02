package main

import (
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/interfaces/cli"
)

func main() {
	engine := core.New()
	go engine.Start()
	cli.StartUICli()
}
