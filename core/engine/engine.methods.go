package engine

import "github.com/robertkrimen/otto"

var vm *otto.Otto

func init() {
	vm = otto.New()
}

// Next Resolves the next logic tree
func (job *Job) next(args []JobArgs) {
	if job.Current == nil {
		return
	}
	nextArgs, err := executeScript(job.Current.script, args)
	if err == nil {
		job.Current = job.Current.positive
	} else {
		job.Current = job.Current.negative
	}
	job.next(nextArgs)
}

// Start Resolves the next logic tree
func (job *Job) Start() {
	job.Current = job.Begin
	job.next(job.Args)
}

// ExecuteScript Executes javascript code
func executeScript(script string, args []JobArgs) ([]JobArgs, error) {
	// Execute script
	vm.Run(script)
	return nil, nil
}
