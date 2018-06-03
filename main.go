package main

import (
	_ "github.com/plopezm/kaiser/config"
	"github.com/plopezm/kaiser/core"
)

func main() {
	engine := core.New()
	engine.Start()
	// cli.StartUICli()
}
