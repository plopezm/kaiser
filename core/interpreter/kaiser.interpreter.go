package interpreter

import (
	httpPlugin "github.com/plopezm/kaiser/plugins/http"
	logPlugin "github.com/plopezm/kaiser/plugins/logger"
	"github.com/robertkrimen/otto"
)

// NewVMWithPlugins Creates a new VM instance using plugins.
// @Param context map[string]interface{} Contains information about the process who creates this VM
func NewVMWithPlugins(context map[string]interface{}) *otto.Otto {
	vm := otto.New()
	registerPlugin(vm, logPlugin.New(context).GetFunctions())
	registerPlugin(vm, httpPlugin.New())
	return vm
}

func registerPlugin(vm *otto.Otto, plugin map[string]interface{}) {
	for key, function := range plugin {
		vm.Set(key, function)
	}
}
