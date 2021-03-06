package types

// JobActivation Represents an activation type
type JobActivation struct {
	Type JobActivationType `json:"type"`
	// Timer represents an ISO 8601 Duration
	Duration string `json:"duration"`
}

// JobActivationType Defines types for launching jobs
type JobActivationType string

const (
	//LOCAL Normal job executed once is received
	LOCAL JobActivationType = "local"
	//REMOTE This job will be executed by request
	REMOTE JobActivationType = "remote"
)
