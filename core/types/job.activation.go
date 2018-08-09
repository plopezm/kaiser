package types

// JobActivation Represents an activation type
type JobActivation struct {
	Type JobActivationType `json:"type"`
	// Timer represents an ISO 8601 Duration
	Duration string    `json:"duration"`
	Args     []JobArgs `json:"args"`
}

// JobActivationType Defines types for launching jobs
type JobActivationType string

const (
	//LOCAL Normal job executed once is received
	LOCAL JobActivationType = "local"
	//GRAPHQL This job will be executed by request
	GRAPHQL JobActivationType = "graphql"
)
