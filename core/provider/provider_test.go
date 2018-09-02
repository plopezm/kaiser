package provider

import (
	"sync"
	"testing"

	"github.com/plopezm/kaiser/core/types"
	"github.com/stretchr/testify/assert"
)

func setup() *types.Job {
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

	return job
}

func TestGetProvider(t *testing.T) {
	// Given

	// When
	provider := GetProvider()

	// Then
	assert := assert.New(t)
	assert.NotNil(provider.Channel)
	assert.NotNil(provider.jobs)
	assert.NotNil(provider.sync)
}

func TestJobProvider_getExistingJobHash_NotFound(t *testing.T) {
	// Given
	provider := GetProvider()

	// When
	bytes, ok := provider.getExistingJobHash("NOT_FOUND")

	// Then
	assert := assert.New(t)
	assert.Nil(bytes)
	assert.Equal(false, ok)
}

func TestJobProvider_addNewJobHash_And_getExistingJobHash(t *testing.T) {
	// Given
	job := setup()
	job.Hash = make([]byte, 1)
	provider := GetProvider()

	// When
	provider.addNewJobHash(job)
	bytes, ok := provider.getExistingJobHash(job.Name)

	// Then
	assert := assert.New(t)
	assert.NotNil(bytes)
	assert.Equal(true, ok)
}

func Test_checkJobHash_SameJob(t *testing.T) {
	// Given
	job := setup()
	job.Hash = make([]byte, 1)
	provider := GetProvider()

	// When
	provider.addNewJobHash(job)

	assert := assert.New(t)
	err := checkJobHash(job)
	t.Log("Job managed")

	// Then
	assert.NotNil(err)
}

func Test_checkJobHash_UpdatingJob(t *testing.T) {
	// Given
	job := setup()
	job.Hash = make([]byte, 1)
	provider := GetProvider()
	var wg sync.WaitGroup

	// When
	provider.addNewJobHash(job)
	job.Hash = make([]byte, 1)
	job.Hash[0] = 1

	assert := assert.New(t)
	wg.Add(1)
	go func() {
		t.Log("Receiving job from provider channel")
		jobReceived := <-provider.Channel
		assert.Equal(job.Name, jobReceived.Name)
		wg.Done()
	}()
	err := checkJobHash(job)
	t.Log("Job managed")

	// Then
	assert.Nil(err)
	wg.Wait()
}
