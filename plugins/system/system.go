package system

import (
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/core/types"
	"github.com/robertkrimen/otto"
)

func init() {
	core.RegisterPlugin(new(OSPlugin))
}

// OSPlugin is used to save process context
type OSPlugin struct {
	context types.JobContext
}

// GetFunctions returns the functions to be registered in the VM
func (plugin *OSPlugin) GetFunctions() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["process"] = map[string]interface{}{
		"sleep": plugin.Sleep,
	}
	functions["system"] = map[string]interface{}{
		"call": plugin.Call,
	}
	return functions
}

// GetInstance Creates a new plugin instance with a context
func (plugin *OSPlugin) GetInstance(context types.JobContext) types.Plugin {
	newPluginInstance := new(OSPlugin)
	newPluginInstance.context = context
	return newPluginInstance
}

// Call Calls an existing job
func (plugin *OSPlugin) Call(jobName string, params map[string]interface{}) otto.Value {
	err := core.GetEngineInstance().ExecuteStoredJob(jobName, params)
	if err != nil {
		res, _ := otto.ToValue(err.Error())
		return res
	}
	return otto.Value{}
}
