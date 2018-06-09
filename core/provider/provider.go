package provider

import (
	"bytes"
	"time"

	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/utils"
)

var provider *JobProvider

func init() {
	// This should prepare everything for thread looking for new files
	provider = new(JobProvider)
	provider.jobs = make(map[string]core.Job)
	provider.Channel = make(chan core.Job)
}

// GetProvider Returns the an instance of a FileJobProvider
func GetProvider() *JobProvider {
	return provider
}

// JobProvider Is a parser who gets the jobs from workspace
type JobProvider struct {
	Channel chan core.Job
	jobs    map[string]core.Job
}

// GetJobs Returns the list of jobs registered
func (provider *JobProvider) GetJobs() []core.Job {
	currentJobs := make([]core.Job, 0)
	for _, job := range provider.jobs {
		currentJobs = append(currentJobs, job)
	}
	return currentJobs
}

// RegisterJobNotifier Create a new listener for other type of notifier
func (provider *JobProvider) RegisterJobNotifier(channel chan core.Job) {
	go observeNotifier(channel)
}

func observeNotifier(notifierChannel chan core.Job) {
	for newJob := range notifierChannel {
		storedJobMetadata, ok := provider.jobs[newJob.Name]

		// If the job has changed, the we should check it
		if !ok || bytes.Compare(storedJobMetadata.Hash, newJob.Hash) != 0 {
			if ok {
				storedJobMetadata.OnDestroy <- true
			}
			newJob.OnDestroy = make(chan bool)

			applyPeriodicity(&newJob)
			provider.jobs[newJob.Name] = newJob
			provider.Channel <- newJob
		}
	}
}

func applyPeriodicity(newJob *core.Job) {
	if len(newJob.Duration) > 0 {
		duration := utils.ParseDuration(newJob.Duration)
		if duration > 0 {
			newJob.Ticker = time.NewTicker(duration)
			go periodHandler(*newJob)
		}
	}
}

func periodHandler(job core.Job) {
	for {
		select {
		case <-job.Ticker.C:
			provider.Channel <- job
		case <-job.OnDestroy:
			job.Ticker.Stop()
			close(job.OnDestroy)
			return
		}
	}
}
