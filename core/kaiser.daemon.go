package core

import (
	"log"
	"sync"
	"time"

	"github.com/plopezm/kaiser/core/engine"
	"github.com/plopezm/kaiser/core/interpreter"
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
	parser      parser.JobParser
	jobs        []engine.Job
	interpreter *interpreter.Interpreter
}

// New Returns the singleton instance of JobEngine
func New() *JobEngine {
	single.Do(func() {
		engineInstance = new(JobEngine)
		//TODO: This part should be implemented using a configuration variable
		engineInstance.parser = file.GetParser()
		engineInstance.parser.Register(&ParserObserver{})
		engineInstance.interpreter = interpreter.New()
	})
	return engineInstance
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for {
		for _, job := range engine.jobs {
			log.Println("[Engine] Executing job:", job.Name)
			job.Start(engine.interpreter)
		}
		time.Sleep(5000 * time.Millisecond)
	}
}
