package interfaces

import (
	"sync"

	"github.com/plopezm/kaiser/core/types"
	"github.com/plopezm/kaiser/core/validation"
)

// Channel the channel used to notify new jobs
var channel chan types.Job
var once sync.Once

func GetChannel() chan types.Job {
	once.Do(func() {
		// This should prepare everything for thread looking for new jobs
		channel = make(chan types.Job)
	})
	return channel
}

// NotifyJob Sends a job to the engine
func NotifyJob(newJob *types.Job) error {
	//newJob.Folder = folder
	//newJob.Hash = hash
	err := validation.VerifyJob(newJob)
	if err != nil {
		return err
	}
	for key, task := range newJob.Tasks {
		task.Name = key
	}
	channel <- *newJob
	return nil
}
