package provider

import (
	"bytes"

	"github.com/plopezm/kaiser/core"
)

var provider *JobProvider

func init() {
	// This should prepare everything for thread looking for new files
	provider = new(JobProvider)
	provider.jobs = make(map[string][]byte)
	provider.Channel = make(chan core.Job)
}

// GetProvider Returns the an instance of a FileJobProvider
func GetProvider() *JobProvider {
	return provider
}

// JobProvider Is a parser who gets the jobs from workspace
type JobProvider struct {
	Channel chan core.Job
	jobs    map[string][]byte
}

// RegisterJobNotifier Create a new listener for other type of notifier
func (provider *JobProvider) RegisterJobNotifier(channel chan core.Job) {
	go observeNotifier(channel)
}

func observeNotifier(notifierChannel chan core.Job) {
	for newJob := range notifierChannel {
		storedJobHash, ok := provider.jobs[newJob.Name]
		// If the job has changed, the we should check it
		if !ok || bytes.Compare(storedJobHash, newJob.Hash) != 0 {
			provider.jobs[newJob.Name] = newJob.Hash
			provider.Channel <- newJob
		}
	}
}
