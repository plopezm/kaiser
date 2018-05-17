package core

import (
	"github.com/plopezm/kaiser/core/interpreter"
	httpPlugin "github.com/plopezm/kaiser/plugins/http"
	logPlugin "github.com/plopezm/kaiser/plugins/logger"
)

func init() {
	interpreter := interpreter.New()
	interpreter.RegisterPlugin(logPlugin.KaiserExports())
	interpreter.RegisterPlugin(httpPlugin.KaiserExports())
}
