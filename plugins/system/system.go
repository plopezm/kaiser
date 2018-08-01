package system

import (
	"github.com/plopezm/kaiser/core/types"
	"github.com/plopezm/kaiser/plugins"
)

// OSPlugin is used to save process context
type OSPlugin struct {
	context types.JobContext
}

// New Creates a new instance of Logger plugin
func New(context types.JobContext) plugins.KaiserPlugin {
	plugin := new(OSPlugin)
	plugin.context = context
	return plugin
}

// GetFunctions returns the functions to be registered in the VM
func (plugin *OSPlugin) GetFunctions() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["system"] = map[string]interface{}{
		"sleep": plugin.Sleep,
	}
	return functions
}
