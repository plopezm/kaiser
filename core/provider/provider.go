package provider

import (
	"bytes"
	"sync"

	"github.com/plopezm/kaiser/core"
)

var provider *JobProvider

func init() {
	// This should prepare everything for thread looking for new files
	provider = new(JobProvider)
	provider.jobs = make(map[string][]byte)
	provider.Channel = make(chan core.Job)
	provider.sync = &sync.Mutex{}
}

// GetProvider Returns the an instance of a FileJobProvider
func GetProvider() *JobProvider {
	return provider
}

// JobProvider Is a parser who gets the jobs from workspace
type JobProvider struct {
	Channel chan core.Job
	jobs    map[string][]byte
	sync    *sync.Mutex
}

// RegisterJobNotifier Create a new listener for other type of notifier
func (provider *JobProvider) RegisterJobNotifier(channel chan core.Job) {
	go observeNotifier(channel)
}

func (provider *JobProvider) addNewJob(newJob *core.Job) {
	provider.sync.Lock()
	defer provider.sync.Unlock()
	provider.jobs[newJob.Name] = newJob.Hash
}

func observeNotifier(notifierChannel chan core.Job) {
	for newJob := range notifierChannel {
		storedJobHash, ok := provider.jobs[newJob.Name]
		// If the job has changed, the we should check it
		if !ok || bytes.Compare(storedJobHash, newJob.Hash) != 0 {
			provider.addNewJob(&newJob)
			provider.Channel <- newJob
		}
	}
}
