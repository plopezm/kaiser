package core

import (
	"github.com/plopezm/kaiser/core/types"
	"github.com/robertkrimen/otto"
)

var plugins []types.Plugin

func init() {
	plugins = make([]types.Plugin, 0)
}

// RegisterPlugin Registers a plugin
func RegisterPlugin(plugin types.Plugin) {
	plugins = append(plugins, plugin)
}

// NewVMWithPlugins Creates a new VM instance using plugins.
// @Param context map[string]interface{} Contains information about the process who creates this VM
func NewVMWithPlugins(context types.JobContext) *otto.Otto {
	vm := otto.New()
	addRegistedPlugins(vm, context)
	return vm
}

func addRegistedPlugins(vm *otto.Otto, context types.JobContext) {
	for _, plugin := range plugins {
		registerPlugin(vm, plugin.GetInstance(context))
	}
}

func registerPlugin(vm *otto.Otto, plugin types.Plugin) {
	for key, function := range plugin.GetFunctions() {
		vm.Set(key, function)
	}
}
