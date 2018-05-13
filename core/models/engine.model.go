package models

// NewJob Returns an initialized job
func NewJob() (job *Job) {
	job = new(Job)
	job.arguments = make(map[string]interface{})
	job.errors = make(map[string]error)
	return job
}

// Job Represents a binary tree
type Job struct {
	Name       string
	Repeatible bool
	current    *DecisionTreeNode
	first      *DecisionTreeNode
	arguments  map[string]interface{}
	isFinished bool
	status     uint8
	errors     map[string]error
}

const (
	// OK Represents that the job finished without errors
	OK = 0
	// RUNNING Represents that the job is still running
	RUNNING = 1
	// ERROR Represents that the job finished with errors
	ERROR = 2
)

// Start Resolves the next logic tree
func (job *Job) Start() {
	job.current = job.first
	job.next(job.arguments)
}

// Next Resolves the next logic tree
func (job *Job) next(args map[string]interface{}) {
	if job.current == nil {
		job.isFinished = true
		return
	}
	nextArgs, err := job.current.runnable.run(args)
	job.errors[job.current.runnable.name()] = err
	if err == nil {
		job.current = job.current.positive
	} else {
		job.current = job.current.negative
	}
	job.next(nextArgs)
}

// DecisionTreeNode Represents a node of a BinaryTree
type DecisionTreeNode struct {
	positive *DecisionTreeNode
	negative *DecisionTreeNode
	runnable Runnable
}

// Runnable Represents a function that can be launched
type Runnable interface {
	name() string
	run(map[string]interface{}) (map[string]interface{}, error)
}
