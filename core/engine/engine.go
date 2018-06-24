package engine

import (
	"sync"

	"github.com/plopezm/kaiser/core/provider"
	"github.com/plopezm/kaiser/core/provider/file"
	"github.com/plopezm/kaiser/core/provider/interfaces"
)

var (
	single sync.Once

	engineInstance *JobEngine
)

// JobEngine Represents the state machine manager
type JobEngine struct {
	provider *provider.JobProvider
}

// New Returns the singleton instance of JobEngine
func New() *JobEngine {
	single.Do(func() {
		engineInstance = new(JobEngine)
		engineInstance.provider = provider.GetProvider()
		engineInstance.provider.RegisterJobNotifier(file.Channel)
		engineInstance.provider.RegisterJobNotifier(interfaces.Channel)
	})
	return engineInstance
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for {
		job := <-engine.provider.Channel
		job.StartNewInstance()
	}
}
