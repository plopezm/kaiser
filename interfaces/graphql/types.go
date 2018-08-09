package graphql

import (
	"encoding/hex"
	"errors"

	graphqlgo "github.com/graphql-go/graphql"
	"github.com/plopezm/kaiser/core/types"
	"github.com/plopezm/kaiser/utils"
)

var (
	jobStatusType = graphqlgo.NewEnum(graphqlgo.EnumConfig{
		Name:        "JobStatus",
		Description: "Status of the current job",
		Values: graphqlgo.EnumValueConfigMap{
			"STOPPED": &graphqlgo.EnumValueConfig{
				Value: types.STOPPED,
			},
			"RUNNING": &graphqlgo.EnumValueConfig{
				Value: types.RUNNING,
			},
		},
	})

	jobArgsType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "jobArgs",
		Description: "Input arguments for all job tasks",
		Fields: graphqlgo.Fields{
			"name": &graphqlgo.Field{
				Type: graphqlgo.NewNonNull(graphqlgo.String),
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if args, ok := p.Source.(types.JobArgs); ok {
						return args.Name, nil
					}
					return nil, errors.New("Error getting JobArgs field " + p.Info.FieldName)
				},
			},
			"value": &graphqlgo.Field{
				Type: graphqlgo.NewNonNull(graphqlgo.String),
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if args, ok := p.Source.(types.JobArgs); ok {
						return args.Value, nil
					}
					return nil, errors.New("Error getting JobArgs field " + p.Info.FieldName)
				},
			},
		},
	})

	jobActivationType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "JobActivation",
		Description: "Job activation options",
		Fields: graphqlgo.Fields{
			"type": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Activation type, currently the options are 'local' or 'graphql'",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if jobActivation, ok := p.Source.(types.JobActivation); ok {
						return jobActivation.Type, nil
					}
					return nil, errors.New("Error getting JobArgs field " + p.Info.FieldName)
				},
			},
			"duration": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "If type is LOCAL, the execution time duration in ISO 8601",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if jobActivation, ok := p.Source.(types.JobActivation); ok {
						return jobActivation.Duration, nil
					}
					return nil, errors.New("Error getting JobArgs field " + p.Info.FieldName)
				},
			},
			"args": &graphqlgo.Field{
				Type:        graphqlgo.NewList(jobArgsType),
				Description: "If type is graphql. Network arguments received with the request",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if jobActivation, ok := p.Source.(types.JobActivation); ok {
						return jobActivation.Args, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
		},
	})

	jobTaskType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "JobTask",
		Description: "Definition of a simple task",
		Fields: graphqlgo.Fields{
			"name": &graphqlgo.Field{
				Type:        graphqlgo.String,
				Description: "The task name",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*types.JobTask); ok {
						return task.Name, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
			"script": &graphqlgo.Field{
				Type:        graphqlgo.String,
				Description: "The script to executed in this task",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*types.JobTask); ok {
						return task.Script, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
			"scriptFile": &graphqlgo.Field{
				Type:        graphqlgo.String,
				Description: "The script file to executed in this task",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*types.JobTask); ok {
						return task.ScriptFile, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
			"onSuccess": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The script to executed if success",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*types.JobTask); ok {
						return task.OnSuccess, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
			"onFailure": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The script to executed if failure",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*types.JobTask); ok {
						return task.OnFailure, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
		},
	})

	//Job definition type for GraphQL
	jobType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "Job",
		Description: "JobType to access using GraphQL",
		Fields: graphqlgo.Fields{
			"version": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The Job engine version, currently only v1 is supported",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						return job.Version, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"name": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The Job name",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						return job.Name, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"args": &graphqlgo.Field{
				Type:        graphqlgo.NewList(jobArgsType),
				Description: "Initial arguments for script tasks. Can be used in the following scripts",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						return job.Args, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"activation": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(jobActivationType),
				Description: "Job activation settings",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						return job.Activation, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"entrypoint": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The first job to be executed",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						return job.Entrypoint, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"tasks": &graphqlgo.Field{
				Type:        graphqlgo.NewList(jobTaskType),
				Description: "The job tasks to be executed",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						taskArray := make([]interface{}, 0)
						for _, jobTask := range job.Tasks {
							taskArray = append(taskArray, jobTask)
						}
						return taskArray, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"status": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(jobStatusType),
				Description: "Current status of the job",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						return job.Status, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"hash": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Job hash file",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						return hex.EncodeToString(job.Hash), nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"log": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Job's output (100 last lines)",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(types.Job); ok {
						content, err := utils.ReadInverseFileContent("logs/"+job.Name+".log", 100)
						return *content, err
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
		},
	})
)
