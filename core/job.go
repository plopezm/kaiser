package core

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/plopezm/kaiser/core/context"
	"github.com/plopezm/kaiser/core/interpreter"
	"github.com/robertkrimen/otto"
)

// InitializeJob Initializes internal attributes of a Job
func InitializeJob(job *Job) {
	job.Status = STOPPED
	job.Sync = &sync.Mutex{}
	job.OnDestroy = make(chan bool)
}

// Job Represents executable job
type Job struct {
	// External attributes
	Version    string              `json:"version"`
	Name       string              `json:"name"`
	Args       []JobArgs           `json:"args"`
	Duration   string              `json:"duration"`
	Entrypoint string              `json:"start"`
	Tasks      map[string]*JobTask `json:"tasks"`
	// Internal attributes
	Status    JobStatus    `json:"status"`
	Sync      *sync.Mutex  `json:"-"`
	Folder    string       `json:"-"`
	Hash      []byte       `json:"-"`
	OnDestroy chan bool    `json:"-"`
	Ticker    *time.Ticker `json:"-"`
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

// Start Resolves the logic tree
func (job *Job) Start() {
	job.Status = RUNNING
	vm := job.initializeVM()
	currentJob := job.Tasks[job.Entrypoint]
	for currentJob != nil {
		switch job.Status {
		case STOPPED:
			return
		case PAUSED:
			job.Sync.Lock()
		default:
		}
		_, err := vm.Run(job.getScript(currentJob))
		if err == nil {
			currentJob = job.Tasks[currentJob.OnSuccess]
		} else {
			currentJob = job.Tasks[currentJob.OnFailure]
		}
	}
	job.Status = STOPPED
}

// Stop Stop job execution
func (job *Job) Stop() {
	job.Status = STOPPED
}

// Pause pauses the process
func (job *Job) Pause() {
	job.Sync.Lock()
	job.Status = PAUSED
}

// Resume unpauses the process
func (job *Job) Resume() {
	job.Status = RUNNING
	job.Sync.Unlock()
}

// initializeVM Creates a new VM with plugins and the current context.
// Every job executed will have its own context, args and plugins.
// By default all plugins are set in the VM.
func (job *Job) initializeVM() *otto.Otto {
	context := &context.JobInstanceContext{
		JobName: job.Name,
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
