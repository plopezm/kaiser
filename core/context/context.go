package context

type JobContext interface {
	GetJobName() string
}

type JobInstanceContext struct {
	JobName string
}

func (context *JobInstanceContext) GetJobName() string {
	return context.JobName
}
