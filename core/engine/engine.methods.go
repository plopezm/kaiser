package engine

import (
	"github.com/robertkrimen/otto"
)

var vm *otto.Otto

func init() {
	vm = otto.New()
}

// Start Resolves the next logic tree
func (job *Job) Start() {
	var nextArgs []JobArgs
	job.Current = job.Tasks[job.Entrypoint]
	for job.Current != nil {
		var err error
		nextArgs, err = executeScript(job.Current.Script, nextArgs)
		if err == nil {
			job.Current = job.Tasks[job.Current.OnSuccess]
		} else {
			job.Current = job.Tasks[job.Current.OnFailure]
		}
	}
}

// ExecuteScript Executes javascript code
func executeScript(script string, args []JobArgs) ([]JobArgs, error) {
	// Execute script
	_, err := vm.Run(script)
	return nil, err
}
