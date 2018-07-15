package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
	"github.com/plopezm/kaiser/core/engine"
)

var jobQuery = graphqlgo.NewObject(graphqlgo.ObjectConfig{
	Name: "jobQuery",
	Fields: graphqlgo.Fields{
		"jobs": &graphqlgo.Field{
			Type: graphqlgo.NewList(jobType), // we return a list of discType, defined above
			Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
				return engine.New().GetJobs(), nil
			},
		},
		"getJobById": &graphqlgo.Field{
			Type: jobType,
			Args: graphqlgo.FieldConfigArgument{
				"jobName": &graphqlgo.ArgumentConfig{
					Description: "Job name",
					Type:        graphqlgo.NewNonNull(graphqlgo.String),
				},
			},
			Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
				var jobName = p.Args["jobName"].(string)
				job, err := engine.New().GetJobByName(jobName)
				return job, err
			},
		},
	},
})
