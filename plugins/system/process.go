package system

import (
	"time"

	"github.com/plopezm/kaiser/core/context"
	"github.com/plopezm/kaiser/plugins"
)

// ProcessPlugin is used to save process context
type ProcessPlugin struct {
	context context.JobContext
}

// New Creates a new instance of Logger plugin
func New(context context.JobContext) plugins.KaiserPlugin {
	plugin := new(ProcessPlugin)
	plugin.context = context
	return plugin
}

// GetFunctions returns the functions to be registered in the VM
func (plugin *ProcessPlugin) GetFunctions() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["system"] = map[string]interface{}{
		"sleep": plugin.Sleep,
	}
	return functions
}

// Sleep Sleeps the current process
func (plugin *ProcessPlugin) Sleep(number int, unit string) {
	var timeUnit time.Duration
	switch unit {
	case "MS":
		timeUnit = time.Millisecond
	case "S":
		timeUnit = time.Second
	default:
		timeUnit = time.Second
	}

	time.Sleep(time.Duration(number) * timeUnit)

}
