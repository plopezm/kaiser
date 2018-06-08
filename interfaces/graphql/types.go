package graphql

import (
	"errors"

	graphqlgo "github.com/graphql-go/graphql"
	"github.com/plopezm/kaiser/core/engine"
)

var (
	jobArgsType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "jobArgs",
		Description: "Input arguments for all job tasks",
		Fields: graphqlgo.Fields{
			"name": &graphqlgo.Field{
				Type: graphqlgo.NewNonNull(graphqlgo.String),
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if args, ok := p.Source.(engine.JobArgs); ok {
						return args.Name, nil
					}
					return nil, errors.New("Error getting JobArgs field " + p.Info.FieldName)
				},
			},
			"value": &graphqlgo.Field{
				Type: graphqlgo.NewNonNull(graphqlgo.String),
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if args, ok := p.Source.(engine.JobArgs); ok {
						return args.Value, nil
					}
					return nil, errors.New("Error getting JobArgs field " + p.Info.FieldName)
				},
			},
		},
	})

	jobTaskType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "JobTask",
		Description: "Definition of a simple task",
		Fields: graphqlgo.Fields{
			"script": &graphqlgo.Field{
				Type:        graphqlgo.String,
				Description: "The script to executed in this task",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*engine.JobTask); ok {
						return task.Script, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
			"scriptFile": &graphqlgo.Field{
				Type:        graphqlgo.String,
				Description: "The script file to executed in this task",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*engine.JobTask); ok {
						return task.ScriptFile, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
			"onSuccess": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The script to executed if success",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*engine.JobTask); ok {
						return task.OnSuccess, nil
					}
					return nil, errors.New("Error getting JobTask field " + p.Info.FieldName)
				},
			},
			"onFailure": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The script to executed if failure",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*engine.JobTask); ok {
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
					if job, ok := p.Source.(engine.Job); ok {
						return job.Version, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"name": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The Job engine version, currently only v1 is supported",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(engine.Job); ok {
						return job.Name, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"args": &graphqlgo.Field{
				Type:        graphqlgo.NewList(jobArgsType),
				Description: "The Job engine version, currently only v1 is supported",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(engine.Job); ok {
						return job.Args, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"duration": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The Job engine version, currently only v1 is supported",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(engine.Job); ok {
						return job.Duration, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"entrypoint": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The Job engine version, currently only v1 is supported",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(engine.Job); ok {
						return job.Entrypoint, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
			"tasks": &graphqlgo.Field{
				Type:        graphqlgo.NewList(jobTaskType),
				Description: "The Job engine version, currently only v1 is supported",
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if job, ok := p.Source.(engine.Job); ok {
						taskArray := make([]interface{}, 0)
						for _, jobTask := range job.Tasks {
							taskArray = append(taskArray, jobTask)
						}
						return taskArray, nil
					}
					return nil, errors.New("Error getting Job field " + p.Info.FieldName)
				},
			},
		},
	})
)
