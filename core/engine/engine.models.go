package engine

// Job Represents a binary tree
type Job struct {
	Version string    `json:"version"`
	Name    string    `json:"name"`
	Args    []JobArgs `json:"args"`
	Begin   *JobTask  `json:"begin"`
	Current *JobTask
}

// JobArgs Represents the input arguments to the executor
type JobArgs struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// JobTask Represents a job task to be performed
type JobTask struct {
	script   string
	positive *JobTask
	negative *JobTask
}
