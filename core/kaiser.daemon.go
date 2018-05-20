package core

import (
	"log"
	"sync"
	"time"

	"github.com/plopezm/kaiser/core/engine"
	"github.com/plopezm/kaiser/core/parser"
	"github.com/plopezm/kaiser/core/parser/file"
)

var (
	single sync.Once

	engineInstance *JobEngine
)

// ParserObserver Type ParseObserver
type ParserObserver struct {
}

// OnNotify Represents a callback when the parser founds new jobs
func (obs ParserObserver) OnNotify(job interface{}) {
	log.Println("Received notification", job)
	engineInstance.jobs = append(engineInstance.jobs, job.(engine.Job))
}

// JobEngine Represents the state machine manager
type JobEngine struct {
	parser parser.JobParser
	jobs   []engine.Job
}

// New Returns the singleton instance of JobEngine
func New() *JobEngine {
	single.Do(func() {
		engineInstance = new(JobEngine)
		//TODO: This part should be implemented using a configuration variable
		engineInstance.parser = file.GetParser()
		engineInstance.parser.Register(&ParserObserver{})
	})
	return engineInstance
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for {
		for i, job := range engine.jobs {
			if job.IsReady() {
				log.Println("[Engine] Executing job:", job.Name)
				go engine.jobs[i].Start()
			}
		}
		time.Sleep(5000 * time.Millisecond)
	}
}
