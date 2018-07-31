package interfaces

import (
	"github.com/plopezm/kaiser/core/types"
)

// Channel the channel used to notify new jobs
var Channel chan types.Job

func init() {
	// This should prepare everything for thread looking for new files
	Channel = make(chan types.Job)
}

// NotifyJob Sends a job to the engine
func NotifyJob(newJob *types.Job) {
	//newJob.Folder = folder
	//newJob.Hash = hash
	for key, task := range newJob.Tasks {
		task.Name = key
	}
	Channel <- *newJob
}
