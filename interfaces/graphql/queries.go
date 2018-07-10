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
	},
})
