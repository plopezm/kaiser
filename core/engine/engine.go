package engine

import (
	"log"
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
	provider    *provider.JobProvider
	jobs        map[string]*core.Job
	jobsMapSync *sync.Mutex
}

// New Returns the singleton instance of JobEngine
func New() *JobEngine {
	single.Do(func() {
		engineInstance = new(JobEngine)
		engineInstance.jobs = make(map[string]*core.Job)
		engineInstance.jobsMapSync = &sync.Mutex{}
		engineInstance.provider = provider.GetProvider()
		engineInstance.provider.RegisterJobNotifier(file.Channel)
		engineInstance.provider.RegisterJobNotifier(interfaces.Channel)
	})
	return engineInstance
}

// GetJobs Returns the list of jobs registered
func (engine *JobEngine) GetJobs() []core.Job {
	engine.jobsMapSync.Lock()
	defer engine.jobsMapSync.Unlock()

	currentJobs := make([]core.Job, 0)
	for _, job := range engine.jobs {
		currentJobs = append(currentJobs, job.Copy())
	}
	return currentJobs
}

func (engine *JobEngine) applyPeriodicity(newJob *core.Job) {
	if len(newJob.Duration) > 0 {
		duration := utils.ParseDuration(newJob.Duration)
		if duration > 0 {
			newJob.Ticker = time.NewTicker(duration)
			go engine.periodHandler(newJob)
		}
	}
}

func (engine *JobEngine) periodHandler(job *core.Job) {
	for {
		select {
		case <-job.Ticker.C:
			engine.executeStoredJob(job.Name)
		case <-job.OnDestroy:
			job.Ticker.Stop()
			close(job.OnDestroy)
			return
		}
	}
}

func (engine *JobEngine) executeStoredJob(jobName string) {
	engine.jobsMapSync.Lock()
	defer engine.jobsMapSync.Unlock()
	storedJob, ok := engine.jobs[jobName]
	if !ok {
		log.Println("Job [" + jobName + "] cannot be executed because it does not exist")
		return
	}
	go storedJob.Start()
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for {
		job := <-engine.provider.Channel
		log.Println("Received job [" + job.Name + "] from provider")
		storedJob, ok := engine.jobs[job.Name]
		if ok {
			storedJob.OnDestroy <- true
		}
		core.InitializeJob(&job)
		engine.applyPeriodicity(&job)
		engine.addJob(&job)

		go job.Start()
	}
}

func (engine *JobEngine) addJob(job *core.Job) {
	engine.jobsMapSync.Lock()
	defer engine.jobsMapSync.Unlock()
	engine.jobs[job.Name] = job
}
