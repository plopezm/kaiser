package main

import (
	"github.com/plopezm/kaiser/core"
)

func main() {
	engine := core.New()
	go engine.Start()
}
