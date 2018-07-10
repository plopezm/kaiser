package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
)

// JobSchema contains job schema, including queries and mutations
var JobSchema, _ = graphqlgo.NewSchema(graphqlgo.SchemaConfig{
	Query:    jobQuery,
	Mutation: jobMutation,
})
