package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
	"github.com/plopezm/kaiser/core/provider/interfaces"
	"github.com/plopezm/kaiser/core/types"
)

var (
	jobArgTypeInput = graphqlgo.NewInputObject(graphqlgo.InputObjectConfig{
		Name: "jobArgTypeInput",
		Fields: graphqlgo.InputObjectConfigFieldMap{
			"name": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Argument name",
			},
			"value": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Argument value",
			},
		},
	})

	jobActivationTypeInput = graphqlgo.NewInputObject(graphqlgo.InputObjectConfig{
		Name: "jobActivationTypeInput",
		Fields: graphqlgo.InputObjectConfigFieldMap{
			"type": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Activation type, currently the options are 'local' or 'graphql'",
			},
			"duration": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "If type is LOCAL, the execution time duration in ISO 8601",
			},
			"args": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewList(jobArgTypeInput),
				Description: "If type is graphql. Network arguments received with the request",
			},
		},
	})

	jobTaskTypeInput = graphqlgo.NewInputObject(graphqlgo.InputObjectConfig{
		Name: "jobTaskTypeInput",
		Fields: graphqlgo.InputObjectConfigFieldMap{
			"name": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "Task name",
			},
			"script": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "script to execute",
			},
			"onSuccess": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "task name to execute if success",
			},
			"onFailure": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "task name to execute if failure",
			},
		},
	})

	createJobType = graphqlgo.NewInputObject(graphqlgo.InputObjectConfig{
		Name: "CreateJobType",
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
				Type:        graphqlgo.NewList(jobArgTypeInput),
				Description: "Initial arguments of the job used in script",
			},
			"activation": &graphqlgo.InputObjectFieldConfig{
				Type:         graphqlgo.NewNonNull(jobActivationTypeInput),
				Description:  "The activation settings",
				DefaultValue: nil,
			},
			"entrypoint": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The first task to be executed",
			},
			"tasks": &graphqlgo.InputObjectFieldConfig{
				Type:        graphqlgo.NewNonNull(graphqlgo.NewList(jobTaskTypeInput)),
				Description: "A list of tasks to perform in this job",
			},
		},
	})

	jobMutation = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name: "jobMutation",
		Fields: graphqlgo.Fields{
			"createJob": &graphqlgo.Field{
				Type: jobType,
				Args: graphqlgo.FieldConfigArgument{
					"input": &graphqlgo.ArgumentConfig{
						Description: "Create Job arguments",
						Type:        graphqlgo.NewNonNull(createJobType),
					},
				},
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					var inp = p.Args["input"].(map[string]interface{})

					newJob := types.Job{
						Name:       inp["name"].(string),
						Entrypoint: inp["entrypoint"].(string),
					}

					newJob.Args = make([]types.JobArgs, len(inp["args"].([]interface{})))
					for index, jobArg := range inp["args"].([]interface{}) {
						newJob.Args[index] = types.JobArgs{
							Name:  jobArg.(map[string]interface{})["name"].(string),
							Value: jobArg.(map[string]interface{})["value"].(string),
						}
					}

					newJob.Tasks = make(map[string]*types.JobTask)
					for _, jobTask := range inp["tasks"].([]interface{}) {
						scriptString := jobTask.(map[string]interface{})["script"].(string)
						newJob.Tasks[jobTask.(map[string]interface{})["name"].(string)] = &types.JobTask{
							Script:    &scriptString,
							OnSuccess: jobTask.(map[string]interface{})["onSuccess"].(string),
							OnFailure: jobTask.(map[string]interface{})["onFailure"].(string),
						}
					}

					var activation = inp["activation"].(map[string]interface{})
					newJob.Activation = types.JobActivation{
						Type:     types.JobActivationType(activation["type"].(string)),
						Duration: activation["duration"].(string),
					}
					newJob.Activation.Args = make([]types.JobArgs, len(activation["args"].([]interface{})))
					for index, jobArg := range activation["args"].([]interface{}) {
						newJob.Args[index] = types.JobArgs{
							Name: jobArg.(map[string]interface{})["name"].(string),
						}
					}

					interfaces.NotifyJob(&newJob)
					return newJob, nil
				},
			},
		},
	})
)
