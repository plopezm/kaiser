package core

import (
	"sync"
	"testing"

	"github.com/plopezm/kaiser/core/types"
	"github.com/stretchr/testify/assert"
)

func TestGetEngineInstance(t *testing.T) {
	// Given
	// When
	engine := GetEngineInstance()
	// Then
	assert := assert.New(t)
	assert.NotNil(engine)
	assert.NotNil(engine.jobs)
	assert.NotNil(engine.jobsMapSync)
	assert.NotNil(engine.provider)
}

func Test_initializeVM(t *testing.T) {
	// Given
	jobName := "TestingJob"
	args := []types.JobArgs{
		{
			Name:  "Var1",
			Value: "Value1",
		},
	}

	// When
	vm := initializeVM(jobName, args)

	// Then
	assert := assert.New(t)
	assert.NotNil(vm)
	valueStored, err := vm.Get("Var1")
	assert.Nil(err)
	assert.Equal("Value1", valueStored.String())
}

func setupJob() *types.Job {
	var job = new(types.Job)
	types.InitializeJob(job)
	script := ""
	scriptFile := "testing.job.json"
	job.Name = "Testing"
	job.Activation = types.JobActivation{
		Type: "remote",
	}
	job.Entrypoint = "testTask"
	job.Tasks = make(map[string]*types.JobTask)
	job.Tasks["testTask"] = &types.JobTask{
		Name:       "testTask",
		OnFailure:  "",
		OnSuccess:  "",
		Script:     &script,
		ScriptFile: &scriptFile,
	}
	job.Parameters = []types.JobArgs{
		{
			Name:  "param1",
			Value: "value1",
		},
	}

	return job
}

func TestJobEngine_addJob(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}
	job := setupJob()

	// When
	engine.addJob(job)

	// Then
	assert := assert.New(t)
	assert.NotNil(engine.jobs)
	assert.Equal(1, len(engine.jobs))
	assert.Equal(job.Name, engine.jobs[job.Name].Name)
}

func TestJobEngine_GetJobs(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}
	job := setupJob()

	// When
	engine.addJob(job)

	// Then
	assert := assert.New(t)
	assert.NotNil(engine.GetJobs())
	assert.Equal(1, len(engine.GetJobs()))
	assert.Equal(job.Name, engine.GetJobs()[0].Name)
}

func TestJobEngine_GetJobByName(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}
	job := setupJob()

	// When
	engine.addJob(job)

	// Then
	assert := assert.New(t)
	storedJob, err := engine.GetJobByName(job.Name)
	assert.Nil(err)
	assert.Equal(job.Name, storedJob.Name)
}

func TestJobEngine_GetJobByName_NotFound(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}

	// When

	// Then
	assert := assert.New(t)
	_, err := engine.GetJobByName("NOT FOUND JOB")
	assert.NotNil(err)
}

func TestJobEngine_prepareJobAndParams_JobNotFound(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}

	// When
	jobPrepared, params, err := engine.prepareJobAndParams("JOB NOT FOUND", nil)

	// Then
	assert := assert.New(t)
	assert.Nil(jobPrepared)
	assert.Nil(params)
	assert.NotNil(err)
}

func TestJobEngine_prepareJobAndParams_NoParameters(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}
	job := setupJob()
	engine.addJob(job)

	// When
	jobPrepared, params, err := engine.prepareJobAndParams(job.Name, nil)

	// Then
	assert := assert.New(t)
	assert.NotNil(jobPrepared)
	assert.NotNil(params)
	assert.Equal(1, len(params))
	assert.Nil(err)
}

func TestJobEngine_prepareJobAndParams_SettingNewParameters(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}
	job := setupJob()
	engine.addJob(job)
	inputParams := make(map[string]interface{})
	inputParams["param1"] = "ChangedByParameter"

	// When
	jobPrepared, params, err := engine.prepareJobAndParams(job.Name, inputParams)

	// Then
	assert := assert.New(t)
	assert.NotNil(jobPrepared)
	assert.NotNil(params)
	assert.Equal(1, len(params))
	assert.Equal("ChangedByParameter", params[0].Value)
	assert.Nil(err)
}

func TestJobEngine_manageActivation(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}
	job := setupJob()
	engine.addJob(job)

	// When
	err := engine.manageActivation(job)

	// Then
	assert := assert.New(t)
	assert.NotNil(err)
}

func TestJobEngine_processReceivedJob(t *testing.T) {
	// Given
	engine := new(JobEngine)
	engine.jobs = make(map[string]*types.Job)
	engine.jobsMapSync = &sync.Mutex{}
	job := setupJob()

	// When
	engineInstance.processReceivedJob(job)

	// Then
	assert := assert.New(t)
	assert.Equal(1, len(engineInstance.jobs))
}
