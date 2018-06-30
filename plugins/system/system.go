package system

import (
	"github.com/plopezm/kaiser/core/context"
	"github.com/plopezm/kaiser/plugins"
)

// OSPlugin is used to save process context
type OSPlugin struct {
	context context.JobContext
}

// New Creates a new instance of Logger plugin
func New(context context.JobContext) plugins.KaiserPlugin {
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
