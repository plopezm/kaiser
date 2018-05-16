package engine

import (
	"log"

	"github.com/plopezm/kaiser/core/engine/models"
	"github.com/plopezm/kaiser/core/interpreter"
)

// Job Represents a binary tree
type Job struct {
	Version    string                     `json:"version"`
	Name       string                     `json:"name"`
	Args       []models.JobArgs           `json:"args"`
	Entrypoint string                     `json:"start"`
	Tasks      map[string]*models.JobTask `json:"tasks"`
	Current    *models.JobTask
}

// Start Resolves the next logic tree
func (job *Job) Start(interpreter *interpreter.Interpreter) {
	var nextArgs = job.Args
	job.Current = job.Tasks[job.Entrypoint]
	for job.Current != nil {
		var err error
		nextArgs, err = interpreter.ExecuteScript(job.Current.Script, nextArgs)
		if err == nil {
			job.Current = job.Tasks[job.Current.OnSuccess]
		} else {
			log.Println(err)
			job.Current = job.Tasks[job.Current.OnFailure]
		}
	}
}
