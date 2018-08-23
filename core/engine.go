package core

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/plopezm/kaiser/core/provider"
	"github.com/plopezm/kaiser/core/provider/file"
	"github.com/plopezm/kaiser/core/provider/interfaces"
	"github.com/plopezm/kaiser/core/types"
	"github.com/plopezm/kaiser/utils"
	"github.com/robertkrimen/otto"
)

var (
	single sync.Once

	engineInstance *JobEngine
)

// JobEngine Represents the state machine manager
type JobEngine struct {
	provider    *provider.JobProvider
	jobs        map[string]*types.Job
	jobsMapSync *sync.Mutex
}

// GetEngineInstance Returns the singleton instance of JobEngine
func GetEngineInstance() *JobEngine {
	single.Do(func() {
		engineInstance = new(JobEngine)
		engineInstance.jobs = make(map[string]*types.Job)
		engineInstance.jobsMapSync = &sync.Mutex{}
		engineInstance.provider = provider.GetProvider()
		engineInstance.provider.RegisterJobNotifier(file.Channel)
		engineInstance.provider.RegisterJobNotifier(interfaces.Channel)
	})
	return engineInstance
}

// initializeVM Creates a new VM with plugins and the current context.
// Every job executed will have its own context, args and plugins.
// By default all plugins are set in the VM.
func initializeVM(jobName string, args []types.JobArgs) *otto.Otto {
	context := &types.JobInstanceContext{
		JobName: jobName,
	}
	vm := NewVMWithPlugins(context)
	// Setting job arguments in VM
	for _, arg := range args {
		vm.Set(arg.Name, arg.Value)
	}
	return vm
}

// GetJobs Returns the list of jobs registered
func (engine *JobEngine) GetJobs() []types.Job {
	engine.jobsMapSync.Lock()
	defer engine.jobsMapSync.Unlock()

	currentJobs := make([]types.Job, 0)
	for _, job := range engine.jobs {
		currentJobs = append(currentJobs, job.Copy())
	}
	return currentJobs
}

// GetJobByName Returns the job with specified name
func (engine *JobEngine) GetJobByName(name string) (types.Job, error) {
	engine.jobsMapSync.Lock()
	defer engine.jobsMapSync.Unlock()

	job, ok := engine.jobs[name]
	if !ok {
		return types.Job{}, errors.New("Job " + name + " not found")
	}
	return job.Copy(), nil
}

func (engine *JobEngine) addJob(job *types.Job) {
	engine.jobsMapSync.Lock()
	defer engine.jobsMapSync.Unlock()
	engine.jobs[job.Name] = job
}

// ExecuteStoredJob Executes an existing job
func (engine *JobEngine) ExecuteStoredJob(jobName string, receivedParameters map[string]types.JobArgs) {
	engine.jobsMapSync.Lock()
	defer engine.jobsMapSync.Unlock()
	storedJob, ok := engine.jobs[jobName]
	if !ok {
		log.Println("Job [" + jobName + "] cannot be executed because it does not exist")
		return
	}

	allParams := make([]types.JobArgs, 0)
	for _, parameter := range storedJob.Parameters {
		if value, ok := receivedParameters[parameter.Name]; ok {
			parameter.Value = value.Value
		}
		allParams = append(allParams, parameter)
	}

	log.Println("-> Executing job [ " + storedJob.Name + " ]")
	go storedJob.Start(initializeVM(storedJob.Name, allParams))
}

func (engine *JobEngine) manageActivation(newJob *types.Job) {
	jobActivation := newJob.Activation
	if jobActivation.Type == types.LOCAL && len(jobActivation.Duration) > 0 {
		duration := utils.ParseDuration(jobActivation.Duration)
		if duration > 0 {
			newJob.Ticker = time.NewTicker(duration)
		}
	}
	go engine.periodHandler(newJob)
}

func (engine *JobEngine) periodHandler(job *types.Job) {

	for {
		select {
		case <-job.Ticker.C:
			engine.ExecuteStoredJob(job.Name, nil)
		case <-job.OnActivation:
			engine.ExecuteStoredJob(job.Name, nil)
		case <-job.OnDestroy:
			job.Clean()
			return
		}
	}
}

// Start Starts engine logic
func (engine *JobEngine) Start() {
	for {
		job := <-engine.provider.Channel
		log.Println("Received job [" + job.Name + "] from provider")
		storedJob, ok := engine.jobs[job.Name]
		if ok {
			storedJob.OnDestroy <- true
		}
		types.InitializeJob(&job)
		engine.manageActivation(&job)
		engine.addJob(&job)
	}
}
