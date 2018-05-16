package core

import (
	"github.com/plopezm/kaiser/core/interpreter"
	jsonPlugin "github.com/plopezm/kaiser/plugins/json"
)

func init() {
	interpreter := interpreter.New()
	interpreter.RegisterPlugin(jsonPlugin.KaiserExports())
}
