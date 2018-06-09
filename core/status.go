package core

type JobStatus int

const (
	// STOPPED The process is stopped
	STOPPED JobStatus = 0
	// RUNNING The process is currently running
	RUNNING JobStatus = 1
)
