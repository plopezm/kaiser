package core

import (
	"sync"

	"github.com/plopezm/kaiser/core/providers/file"
)

var (
	single sync.Once

	engineInstance *JobEngine
)

// JobEngine Represents the state machine manager
type JobEngine struct {
	fileJobProvider *file.JobProvider
}

// New Returns the singleton instance of JobEngine
func New() *JobEngine {
	single.Do(func() {
		engineInstance = new(JobEngine)
		engineInstance.fileJobProvider = file.GetProvider()
	})
	return engineInstance
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for {
		job := <-engine.fileJobProvider.Channel
		go job.Start()
	}
}
