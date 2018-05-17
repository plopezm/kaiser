package interpreter

import (
	"sync"

	EngineModels "github.com/plopezm/kaiser/core/engine/models"
	"github.com/robertkrimen/otto"
)

var once sync.Once
var interpreter *Interpreter

func init() {
}

func New() *Interpreter {
	once.Do(func() {
		interpreter = new(Interpreter)
		interpreter.VM = otto.New()
	})
	return interpreter
}

type Interpreter struct {
	VM *otto.Otto
}

func (interpreter *Interpreter) RegisterPlugin(plugin map[string]interface{}) {
	for key, function := range plugin {
		interpreter.VM.Set(key, function)
	}
}

func (interpreter *Interpreter) ExecuteScript(script string, args []EngineModels.JobArgs) ([]EngineModels.JobArgs, error) {
	// Execute script
	for _, arg := range args {
		interpreter.VM.Set(arg.Name, arg.Value)
	}
	_, err := interpreter.VM.Run(script)
	return nil, err
}
