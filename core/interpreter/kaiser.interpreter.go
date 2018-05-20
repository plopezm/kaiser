package interpreter

import (
	httpPlugin "github.com/plopezm/kaiser/plugins/http"
	logPlugin "github.com/plopezm/kaiser/plugins/logger"
	"github.com/robertkrimen/otto"
)

func NewVM() *otto.Otto {
	vm := otto.New()
	return vm
}

func NewVMWithPlugins() *otto.Otto {
	vm := otto.New()
	registerPlugin(vm, logPlugin.KaiserExports())
	registerPlugin(vm, httpPlugin.KaiserExports())
	return vm
}

func registerPlugin(vm *otto.Otto, plugin map[string]interface{}) {
	for key, function := range plugin {
		vm.Set(key, function)
	}
}
