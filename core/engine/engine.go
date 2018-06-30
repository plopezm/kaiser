package engine

import (
	"sync"
	"time"

	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/core/provider"
	"github.com/plopezm/kaiser/core/provider/file"
	"github.com/plopezm/kaiser/core/provider/interfaces"
	"github.com/plopezm/kaiser/utils"
)

var (
	single sync.Once

	engineInstance *JobEngine
)

// JobEngine Represents the state machine manager
type JobEngine struct {
	provider *provider.JobProvider
	jobs     map[string]*core.Job
}

// New Returns the singleton instance of JobEngine
func New() *JobEngine {
	single.Do(func() {
		engineInstance = new(JobEngine)
		engineInstance.jobs = make(map[string]*core.Job)
		engineInstance.provider = provider.GetProvider()
		engineInstance.provider.RegisterJobNotifier(file.Channel)
		engineInstance.provider.RegisterJobNotifier(interfaces.Channel)
	})
	return engineInstance
}

// GetJobs Returns the list of jobs registered
func (engine *JobEngine) GetJobs() []core.Job {
	currentJobs := make([]core.Job, 0)
	for _, job := range engine.jobs {
		jobCopy := *job
		currentJobs = append(currentJobs, jobCopy)
	}
	return currentJobs
}

func (engine *JobEngine) applyPeriodicity(newJob *core.Job) {
	if len(newJob.Duration) > 0 {
		duration := utils.ParseDuration(newJob.Duration)
		if duration > 0 {
			newJob.Ticker = time.NewTicker(duration)
			go engine.periodHandler(*newJob)
		}
	}
}

func (engine *JobEngine) periodHandler(job core.Job) {
	for {
		select {
		case <-job.Ticker.C:
			engine.provider.Channel <- job
		case <-job.OnDestroy:
			job.Ticker.Stop()
			close(job.OnDestroy)
			return
		}
	}
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for {
		job := <-engine.provider.Channel

		storedJob, ok := engine.jobs[job.Name]
		if ok {
			storedJob.OnDestroy <- true
		}
		core.InitializeJob(&job)
		engine.applyPeriodicity(&job)
		engine.jobs[job.Name] = &job

		go job.Start()
	}
}
