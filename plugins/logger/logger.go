package logger

import (
	"log"
	"os"

	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/core/types"
)

func init() {
	core.RegisterPlugin(new(LogPlugin))
}

// LogPlugin is used to save process context
type LogPlugin struct {
	context types.JobContext
}

// GetInstance Creates a new plugin instance with a context
func (plugin *LogPlugin) GetInstance(context types.JobContext) types.Plugin {
	newPluginInstance := new(LogPlugin)
	newPluginInstance.context = context
	return newPluginInstance
}

// GetFunctions returns the functions to be registered in the VM
func (plugin *LogPlugin) GetFunctions() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["logger"] = map[string]interface{}{
		"info": plugin.Info,
	}
	return functions
}

// Info Prints objects or strings sent as parameters
func (plugin *LogPlugin) Info(args ...interface{}) {
	f, err := os.OpenFile("logs/"+plugin.context.GetJobName()+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Error creating file logs/" + plugin.context.GetJobName() + ".log")
		return
	}
	defer f.Close()
	logger := log.New(f, "", log.Ldate|log.Ltime|log.LUTC)
	for _, arg := range args {
		logger.Println(arg)
	}
}
