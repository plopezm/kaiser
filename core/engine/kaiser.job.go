package engine

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/plopezm/kaiser/core/interpreter"
	"github.com/robertkrimen/otto"
)

// Job Represents executable job
type Job struct {
	Version    string              `json:"version"`
	Name       string              `json:"name"`
	Args       []JobArgs           `json:"args"`
	Repeatable bool                `json:"repeatable"`
	Entrypoint string              `json:"start"`
	Tasks      map[string]*JobTask `json:"tasks"`
	current    *JobTask
	executed   bool
	Folder     string
	VM         *otto.Otto
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

// IsReady Returns true if a job is ready to be executed
func (job *Job) IsReady() bool {
	return (!job.Repeatable && !job.executed) || (job.Repeatable)
}

// Start Resolves the next logic tree
func (job *Job) Start() {
	if job.executed && !job.Repeatable {
		return
	}
	job.executed = true

	job.VM = interpreter.NewVMWithPlugins()

	job.setArguments(job.Args...)
	job.current = job.Tasks[job.Entrypoint]
	for job.current != nil {
		var err error
		err = job.executeTask()
		if err == nil {
			job.current = job.Tasks[job.current.OnSuccess]
		} else {
			job.current = job.Tasks[job.current.OnFailure]
		}
	}
}

func (job *Job) setArguments(args ...JobArgs) {
	for _, arg := range job.Args {
		job.VM.Set(arg.Name, arg.Value)
	}
}

// ExecuteTask Executes the current task
func (job *Job) executeTask() error {
	_, err := job.VM.Run(job.getScript())
	return err
}

// GetScript Returns the script from inline declaration or from referenced declaration
func (job *Job) getScript() string {
	if job.current.Script != nil {
		return *job.current.Script
	}
	if job.current.ScriptFile != nil {
		raw, err := ioutil.ReadFile(job.Folder + *job.current.ScriptFile)
		if err != nil {
			log.Fatalln(err.Error())
			os.Exit(1)
		}
		return string(raw)
	}
	return "console.log('[VM ERROR]: Error getScript(), maybe script and scriptFile are undefined')"
}
