package interpreter

import (
	"sync"

	EngineModels "github.com/plopezm/kaiser/core/engine/models"
	"github.com/robertkrimen/otto"
)

var vm *otto.Otto
var once sync.Once
var interpreter *Interpreter

func init() {
	vm = otto.New()
}

func New() *Interpreter {
	once.Do(func() {
		interpreter = new(Interpreter)
	})
	return interpreter
}

type Interpreter struct {
}

func (interpreter *Interpreter) RegisterPlugin(plugin map[string]interface{}) {
	for key, function := range plugin {
		vm.Set(key, function)
	}
}

func (interpreter *Interpreter) ExecuteScript(script string, args []EngineModels.JobArgs) ([]EngineModels.JobArgs, error) {
	// Execute script
	for _, arg := range args {
		vm.Set(arg.Name, arg.Value)
	}
	_, err := vm.Run(script)
	return nil, err
}
