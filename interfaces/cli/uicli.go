package cli

import (
	"time"

	"github.com/plopezm/kaiser/core/provider"
)

var jobProvider *provider.JobProvider

const (
	JOB_LIST_COMMAND = "job list"
	EXIT_COMMAND     = "exit"
)

// StartUICli Starts CLI UI
func StartUICli() {
	jobProvider = provider.GetProvider()
	for {
		time.Sleep(5000 * time.Millisecond)
	}
}
