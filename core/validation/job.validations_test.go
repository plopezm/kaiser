package validation

import (
	"testing"

	"github.com/plopezm/kaiser/core/types"
	"github.com/storozhukBM/verifier"
	"github.com/stretchr/testify/assert"
)

func Test_verifyTask_True(t *testing.T) {
	// Given
	verify := verifier.New()
	script := "Something;"
	task := &types.JobTask{
		Script: &script,
	}
	// When
	verifyTask(task, verify)
	// Then
	assert.Nil(t, verify.GetError())
}

func Test_verifyTask_Nil_Script_False(t *testing.T) {
	// Given
	verify := verifier.New()
	task := &types.JobTask{
		Script: nil,
	}
	// When
	verifyTask(task, verify)
	// Then
	assert.NotNil(t, verify.GetError())
}

func Test_verifyTasks_True(t *testing.T) {
	// Given
	verify := verifier.New()
	script := "Something;"
	tasks := map[string]*types.JobTask{
		"task1": &types.JobTask{
			Script: &script,
		},
	}
	// When
	verifyTasks(tasks, verify)
	// Then
	assert.Nil(t, verify.GetError())
}

func Test_verifyTasks_False(t *testing.T) {
	// Given
	verify := verifier.New()
	tasks := map[string]*types.JobTask{}
	// When
	verifyTasks(tasks, verify)
	// Then
	assert.NotNil(t, verify.GetError())
}

func Test_verifyActivation_True(t *testing.T) {
	// Given
	verify := verifier.New()
	activation := types.JobActivation{
		Type:     types.REMOTE,
		Duration: "",
	}

	// When
	verifyActivation(activation, verify)
	// Then
	assert.Nil(t, verify.GetError())
}

func Test_verifyActivation_Local_No_Duration_False(t *testing.T) {
	// Given
	verify := verifier.New()
	activation := types.JobActivation{
		Type:     types.LOCAL,
		Duration: "",
	}

	// When
	verifyActivation(activation, verify)
	// Then
	assert.NotNil(t, verify.GetError())
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

func TestVerifyJob_True(t *testing.T) {
	// Given
	job := setupJob()
	// When
	err := VerifyJob(job)
	// Then
	assert.Nil(t, err)
}

func TestVerifyJob_False(t *testing.T) {
	// Given
	job := setupJob()
	job.Name = ""
	// When
	err := VerifyJob(job)
	// Then
	assert.NotNil(t, err)
}
