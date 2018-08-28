package validation

import (
	"github.com/plopezm/kaiser/core/types"
	"github.com/storozhukBM/verifier"
)

// VerifyJob Validates Job fields
func VerifyJob(newJob *types.Job) error {
	verify := verifier.New()
	verify.That(newJob != nil, "The job cannot be null")
	verify.That(len(newJob.Name) > 0, "The name cannot be empty")
	verifyActivation(newJob.Activation, verify)
	verifyTasks(newJob.Tasks, verify)
	verify.That(newJob.TaskExist(newJob.Entrypoint), "The entrypoint task should exist in the task list")

	return verify.GetError()
}

func verifyActivation(activation types.JobActivation, verify *verifier.Verify) {
	if activation.Type == types.LOCAL {
		verify.That(len(activation.Duration) > 0, "Duration didn't set")
	}
}

func verifyTasks(tasks map[string]*types.JobTask, verify *verifier.Verify) {
	verify.That(tasks != nil, "Task map should be defined")
	verify.That(len(tasks) > 0, "There should be one task al least")
	for name, task := range tasks {
		verify.That(len(name) > 0, "The task name cannot be empty")
		verifyTask(task, verify)
	}
}

func verifyTask(task *types.JobTask, verify *verifier.Verify) {
	verify.That(task.Script != nil, "Script should be defined")
}
