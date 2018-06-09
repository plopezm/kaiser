package logger

import (
	"log"
	"os"

	"github.com/plopezm/kaiser/core/context"
	"github.com/plopezm/kaiser/plugins"
)

// LogPlugin is used to save process context
type LogPlugin struct {
	context context.JobContext
}

// New Creates a new instance of Logger plugin
func New(context context.JobContext) plugins.KaiserPlugin {
	plugin := new(LogPlugin)
	plugin.context = context
	return plugin
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
