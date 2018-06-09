package interpreter

import (
	"github.com/plopezm/kaiser/core/context"
	"github.com/plopezm/kaiser/plugins"
	httpPlugin "github.com/plopezm/kaiser/plugins/http"
	logPlugin "github.com/plopezm/kaiser/plugins/logger"
	"github.com/robertkrimen/otto"
)

// NewVMWithPlugins Creates a new VM instance using plugins.
// @Param context map[string]interface{} Contains information about the process who creates this VM
func NewVMWithPlugins(context context.JobContext) *otto.Otto {
	vm := otto.New()
	registerPlugin(vm, logPlugin.New(context))
	registerPlugin(vm, httpPlugin.New(context))
	return vm
}

func registerPlugin(vm *otto.Otto, plugin plugins.KaiserPlugin) {
	for key, function := range plugin.GetFunctions() {
		vm.Set(key, function)
	}
}
