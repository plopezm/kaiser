package interfaces

import (
	"os"
	"sync"
	"testing"

	"github.com/plopezm/kaiser/core/types"
	"github.com/stretchr/testify/assert"
)

func Test_GetChannel(t *testing.T) {
	// When
	ch := GetChannel()

	// Then
	assert := assert.New(t)
	assert.NotNil(ch)
}

func setup() *types.Job {
	os.Mkdir("workspace", 0700)
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

func TestNotifyJob_Notified(t *testing.T) {
	// Given
	ch := GetChannel()
	var wg sync.WaitGroup
	job := setup()

	// When
	wg.Add(1)
	go func() {
		t.Log("Starting reading from channel")
		job := <-ch
		t.Log("Job read from channel")
		assert.Equal(t, "Testing", job.Name)
		wg.Done()
	}()
	err := NotifyJob(job)
	t.Log("Job parsed")

	// Then
	assert := assert.New(t)
	assert.Nil(err)
	wg.Wait()
}

func TestNotifyJob_Invalid(t *testing.T) {
	// Given
	job := setup()
	job.Entrypoint = ""

	// When
	err := NotifyJob(job)
	t.Log("Job parsed")

	// Then
	assert := assert.New(t)
	assert.NotNil(err)
}
