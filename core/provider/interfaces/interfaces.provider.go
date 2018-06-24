package interfaces

import (
	"github.com/plopezm/kaiser/core"
)

var Channel chan core.Job

func init() {
	// This should prepare everything for thread looking for new files
	Channel = make(chan core.Job)
}

// NotifyJob Sends a job to the engine
func NotifyJob(newJob *core.Job) {
	//newJob.Folder = folder
	//newJob.Hash = hash
	Channel <- *newJob
}
