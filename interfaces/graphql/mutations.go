package graphql

import (
	"fmt"

	graphqlgo "github.com/graphql-go/graphql"
	"github.com/plopezm/kaiser/core"
)

var (
	createJobType = graphqlgo.NewInputObject(graphqlgo.InputObjectConfig{
		Name: "CreateJob",
		Fields: graphqlgo.InputObjectConfigFieldMap{
			"version": &graphqlgo.InputObjectFieldConfig{
				Type:         graphqlgo.String,
				Description:  "Job version",
				DefaultValue: "1",
			},
			"name": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Job name",
			},
			"args": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.Int),
				Description: "Initial arguments of the job used in script",
			},
			"duration": &graphqlgo.InputObjectFieldConfig{
				Type:         graphqlgo.String,
				Description:  "The period of time between executions",
				DefaultValue: nil,
			},
			"entrypoint": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The first task to be executed",
			},
			"task": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.NewList(jobTaskType)),
				Description: "A list of tasks to perform in this job",
			},
		},
	})

	mutationType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name: "MutationType",
		Fields: graphqlgo.Fields{
			"createJobType": &graphqlgo.Field{
				Type: jobType,
				Args: graphqlgo.FieldConfigArgument{
					"input": &graphqlgo.ArgumentConfig{
						Description: "Create Job arguments",
						Type:        graphqlgo.NewNonNull(createJobType),
					},
				},
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					var inp = p.Args["input"].(map[string]interface{})

					newJob := core.Job{
						Name:       inp["name"].(string),
						Args:       inp["args"].([]core.JobArgs),
						Entrypoint: inp["entrypoint"].(string),
						Duration:   inp["duration"].(string),
						Tasks:      inp["tasks"].(map[string]*core.JobTask),
					}
					fmt.Println(newJob)
					return newJob, nil
				},
			},
		},
	})
)
