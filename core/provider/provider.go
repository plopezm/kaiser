package provider

import (
	"bytes"
	"sync"

	"github.com/plopezm/kaiser/core/types"
)

var (
	single sync.Once

	provider *JobProvider
)

// GetProvider Returns the an instance of a FileJobProvider
func GetProvider() *JobProvider {
	single.Do(func() {
		provider = new(JobProvider)
		provider.jobs = make(map[string][]byte)
		provider.Channel = make(chan types.Job)
		provider.sync = &sync.Mutex{}
	})
	return provider
}

// JobProvider Is a parser who gets the jobs from workspace
type JobProvider struct {
	Channel chan types.Job
	jobs    map[string][]byte
	sync    *sync.Mutex
}

// RegisterJobNotifier Create a new listener for other type of notifier
func (provider *JobProvider) RegisterJobNotifier(channel chan types.Job) {
	go observeNotifier(channel)
}

func (provider *JobProvider) addNewJobHash(newJob *types.Job) {
	provider.sync.Lock()
	defer provider.sync.Unlock()
	provider.jobs[newJob.Name] = newJob.Hash
}

func (provider *JobProvider) getExistingJobHash(name string) ([]byte, bool) {
	provider.sync.Lock()
	defer provider.sync.Unlock()
	storedJobHash, ok := provider.jobs[name]
	return storedJobHash, ok
}

func observeNotifier(notifierChannel chan types.Job) {
	for newJob := range notifierChannel {
		storedJobHash, ok := provider.getExistingJobHash(newJob.Name)
		// If the job has changed, the we should check it
		if !ok || bytes.Compare(storedJobHash, newJob.Hash) != 0 {
			provider.addNewJobHash(&newJob)
			provider.Channel <- newJob
		}
	}
}
