package core

import (
	"github.com/plopezm/kaiser/core/models"
	"github.com/plopezm/kaiser/utils/observer"
)

// JobParser Represents the minimal functions that a parser should provide
type JobParser interface {
	Register(observer.Observer)
	Deregister(observer.Observer)
	Notify(observer.Event)
	GetJobs() map[string]models.Job
}
