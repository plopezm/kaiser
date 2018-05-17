package logger

import (
	"log"
)

// KaiserExports Used by Kaiser, returns new functionality for Kaiser
func KaiserExports() (functions map[string]interface{}) {
	functions = make(map[string]interface{})
	functions["logger"] = map[string]interface{}{
		"info": Info,
	}
	return functions
}

// Info Prints objects or strings sent as parameters
func Info(args ...interface{}) {
	for _, arg := range args {
		log.Println(arg)
	}
}
