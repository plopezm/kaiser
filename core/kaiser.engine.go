package core

import (
	"log"
	"sync"
	"time"

	"github.com/plopezm/kaiser/core/parsers/file"
	"github.com/plopezm/kaiser/utils/observer"
)

var (
	single sync.Once

	engine *JobEngine
)

type ParserObserver struct {
}

func (obs ParserObserver) OnNotify(e observer.Event) {
	log.Println("Received notification", e)
}

// JobEngine Represents the state machine manager
type JobEngine struct {
	parser JobParser
	jobs   []engine.Job
}

// New Returns the singleton instance of JobEngine
func New() *JobEngine {
	single.Do(func() {
		engine = new(JobEngine)
		//TODO: This part should be implemented using a configuration variable
		engine.parser = file.GetParser()
		engine.parser.Register(&ParserObserver{})
	})
	return engine
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for job := range engine.jobs {
		job.Start()
		time.Sleep(1000 * time.Millisecond)
	}
}
