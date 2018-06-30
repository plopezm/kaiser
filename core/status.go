package core

// JobStatus of the current process
type JobStatus int

const (
	// STOPPED The process is stopped
	STOPPED JobStatus = 0
	// RUNNING The process is currently running
	RUNNING JobStatus = 1
	// PAUSED The process is currently paused
	PAUSED JobStatus = 2
)
