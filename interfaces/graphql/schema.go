package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
	"github.com/plopezm/kaiser/core/provider"
)

var queryType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
	Name: "Query",
	Fields: graphqlgo.Fields{
		"jobs": &graphqlgo.Field{
			Type: graphqlgo.NewList(jobType), // we return a list of discType, defined above
			Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
				return provider.GetProvider().GetJobs(), nil
			},
		},
	},
})

var JobSchema, _ = graphqlgo.NewSchema(graphqlgo.SchemaConfig{
	Query: queryType,
	// mutation will be added later
})
