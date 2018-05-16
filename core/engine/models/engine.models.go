package models

// JobArgs Represents the input arguments to the executor
type JobArgs struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// JobTask Represents a job task to be performed
type JobTask struct {
	Script    string `json:"script"`
	OnSuccess string `json:"onSuccess"`
	OnFailure string `json:"onFailure"`
}
