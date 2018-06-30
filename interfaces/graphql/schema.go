package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
	"github.com/plopezm/kaiser/core/engine"
)

var queryType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
	Name: "Query",
	Fields: graphqlgo.Fields{
		"jobs": &graphqlgo.Field{
			Type: graphqlgo.NewList(jobType), // we return a list of discType, defined above
			Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
				return engine.New().GetJobs(), nil
			},
		},
	},
})

var JobSchema, _ = graphqlgo.NewSchema(graphqlgo.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})
