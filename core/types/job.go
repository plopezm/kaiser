package types

import (
	"sync"
	"time"

	"github.com/robertkrimen/otto"
)

// InitializeJob Initializes internal attributes of a Job
func InitializeJob(job *Job) {
	job.Status = STOPPED
	job.sync = &sync.Mutex{}
	job.statusSync = &sync.Mutex{}
	job.OnActivation = make(chan bool)
	job.OnDestroy = make(chan bool)
	job.Ticker = &time.Ticker{
		C: nil,
	}
}

// Job Represents executable job
type Job struct {
	// External attributes
	Version    string              `json:"version"`
	Name       string              `json:"name"`
	Args       []JobArgs           `json:"args"`
	Activation JobActivation       `json:"activation"`
	Entrypoint string              `json:"start"`
	Tasks      map[string]*JobTask `json:"tasks"`
	// Internal attributes
	sync         *sync.Mutex
	Status       JobStatus `json:"status"`
	statusSync   *sync.Mutex
	Folder       string       `json:"-"`
	Hash         []byte       `json:"hash"`
	OnActivation chan bool    `json:"-"`
	OnDestroy    chan bool    `json:"-"`
	Ticker       *time.Ticker `json:"-"`
}

// JobArgs Represents the input arguments to the executor
type JobArgs struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// JobTask Represents a job task to be performed
type JobTask struct {
	Name       string  `json:"name"`
	Script     *string `json:"script"`
	ScriptFile *string `json:"scriptFile"`
	OnSuccess  string  `json:"onSuccess"`
	OnFailure  string  `json:"onFailure"`
}

// Start Resolves the logic tree
func (job *Job) Start(vm *otto.Otto) {
	job.sync.Lock()
	defer job.sync.Unlock()
	job.SetStatus(RUNNING)
	currentJob := job.Tasks[job.Entrypoint]
	for currentJob != nil {
		switch job.Status {
		case STOPPED:
			return
		default:
		}
		_, err := vm.Run(job.getScript(currentJob))
		if err == nil {
			currentJob = job.Tasks[currentJob.OnSuccess]
		} else {
			currentJob = job.Tasks[currentJob.OnFailure]
		}
	}
	job.SetStatus(STOPPED)
}

// SetStatus Sets the job status in an atomic way
func (job *Job) SetStatus(status JobStatus) {
	job.statusSync.Lock()
	defer job.statusSync.Unlock()
	job.Status = status
}

// GetStatus Returns the job status in an atomic way
func (job *Job) GetStatus() JobStatus {
	job.statusSync.Lock()
	defer job.statusSync.Unlock()
	return job.Status
}

// Stop Stop job execution
func (job *Job) Stop() {
	job.SetStatus(STOPPED)
}

// Copy creates a copy of a job object
func (job *Job) Copy() (copy Job) {
	copy.Version = job.Version
	copy.Name = job.Name
	copy.Status = job.GetStatus()
	copy.Hash = job.Hash
	copy.Tasks = job.Tasks
	copy.Args = job.Args
	copy.Activation = job.Activation
	copy.Entrypoint = job.Entrypoint
	return copy
}

// Clean Removes reserved channels and timers
func (job *Job) Clean() {
	job.Ticker.Stop()
	close(job.OnDestroy)
	close(job.OnActivation)
}

// GetScript Returns the script from inline declaration or from referenced declaration
func (job *Job) getScript(current *JobTask) string {
	if current.Script != nil {
		return *current.Script
	}
	return "console.log('[VM ERROR]: Error getScript(), maybe script and scriptFile are undefined')"
}
