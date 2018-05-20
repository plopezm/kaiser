package engine

import (
	"log"

	"github.com/plopezm/kaiser/core/engine/task"
	"github.com/plopezm/kaiser/core/interpreter"
)

// Job Represents a binary tree
type Job struct {
	Version    string                   `json:"version"`
	Name       string                   `json:"name"`
	Args       []task.JobArgs           `json:"args"`
	Repeatable bool                     `json:"repeatable"`
	Entrypoint string                   `json:"start"`
	Tasks      map[string]*task.JobTask `json:"tasks"`
	current    *task.JobTask
	executed   bool
	Folder     string
}

func (job *Job) IsReady() bool {
	return (!job.Repeatable && !job.executed) || (job.Repeatable)
}

// Start Resolves the next logic tree
func (job *Job) Start(interpreter *interpreter.Interpreter) {
	if job.executed && !job.Repeatable {
		return
	}
	job.executed = true

	var nextArgs = job.Args
	job.current = job.Tasks[job.Entrypoint]
	for job.current != nil {
		var err error
		nextArgs, err = interpreter.ExecuteScript(job.current.GetScript(job.Folder), nextArgs)
		if err == nil {
			job.current = job.Tasks[job.current.OnSuccess]
		} else {
			log.Println(err)
			job.current = job.Tasks[job.current.OnFailure]
		}
	}
}
