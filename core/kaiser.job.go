package core

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/plopezm/kaiser/core/interpreter"
	"github.com/robertkrimen/otto"
)

// Job Represents executable job
type Job struct {
	Version    string                 `json:"version"`
	Name       string                 `json:"name"`
	Args       []JobArgs              `json:"args"`
	Duration   string                 `json:"duration"`
	Entrypoint string                 `json:"start"`
	Tasks      map[string]*JobTask    `json:"tasks"`
	Folder     string                 `json:"-"`
	Hash       []byte                 `json:"-"`
	OnDestroy  chan bool              `json:"-"`
	Ticker     *time.Ticker           `json:"-"`
	Instances  map[string]JobInstance `json:"-"`
}

// JobArgs Represents the input arguments to the executor
type JobArgs struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// JobTask Represents a job task to be performed
type JobTask struct {
	Script     *string `json:"script"`
	ScriptFile *string `json:"scriptFile"`
	OnSuccess  string  `json:"onSuccess"`
	OnFailure  string  `json:"onFailure"`
}

type JobInstance struct {
	status JobInstanceStatus
}

type JobInstanceStatus int

const (
	STOPPED JobInstanceStatus = 0
	RUNNING JobInstanceStatus = 1
)

// StartNewInstance Resolves the next logic tree
func (job *Job) StartNewInstance() {
	log.Println("Running job: " + job.Name)
	go func() {
		vm := job.initializeVM()
		currentJob := job.Tasks[job.Entrypoint]
		for currentJob != nil {
			_, err := vm.Run(job.getScript(currentJob))
			if err == nil {
				currentJob = job.Tasks[currentJob.OnSuccess]
			} else {
				currentJob = job.Tasks[currentJob.OnFailure]
			}
		}
	}()
}

// initializeVM Creates a new VM with plugins and the current context.
// Every job executed will have its own context, args and plugins.
// By default all plugins are set in the VM.
func (job *Job) initializeVM() *otto.Otto {
	context := map[string]interface{}{
		"jobName": job.Name,
	}
	vm := interpreter.NewVMWithPlugins(context)
	// Setting job arguments in VM
	for _, arg := range job.Args {
		vm.Set(arg.Name, arg.Value)
	}
	return vm
}

// GetScript Returns the script from inline declaration or from referenced declaration
func (job *Job) getScript(current *JobTask) string {
	if current.Script != nil {
		return *current.Script
	}
	if current.ScriptFile != nil {
		raw, err := ioutil.ReadFile(job.Folder + *current.ScriptFile)
		if err != nil {
			log.Fatalln(err.Error())
			os.Exit(1)
		}
		return string(raw)
	}
	return "console.log('[VM ERROR]: Error getScript(), maybe script and scriptFile are undefined')"
}
