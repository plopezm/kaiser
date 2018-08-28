package interfaces

import (
	"github.com/plopezm/kaiser/core/types"
	"github.com/plopezm/kaiser/core/validation"
)

// Channel the channel used to notify new jobs
var Channel chan types.Job

func init() {
	// This should prepare everything for thread looking for new files
	Channel = make(chan types.Job)
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
	Channel <- *newJob
	return nil
}
