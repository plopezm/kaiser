package core

import (
	"github.com/plopezm/kaiser/core/types"
	"github.com/plopezm/kaiser/plugins"
	httpPlugin "github.com/plopezm/kaiser/plugins/http"
	logPlugin "github.com/plopezm/kaiser/plugins/logger"
	systemPlugin "github.com/plopezm/kaiser/plugins/system"
	"github.com/robertkrimen/otto"
)

// NewVMWithPlugins Creates a new VM instance using plugins.
// @Param context map[string]interface{} Contains information about the process who creates this VM
func NewVMWithPlugins(context types.JobContext) *otto.Otto {
	vm := otto.New()
	registerPlugin(vm, logPlugin.New(context))
	registerPlugin(vm, httpPlugin.New(context))
	registerPlugin(vm, systemPlugin.New(context))
	return vm
}

func registerPlugin(vm *otto.Otto, plugin plugins.KaiserPlugin) {
	for key, function := range plugin.GetFunctions() {
		vm.Set(key, function)
	}
}