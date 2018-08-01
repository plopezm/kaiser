package graphql

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/plopezm/kaiser/core"
)

func init() {
	graphqlHandler := handler.New(&handler.Config{
		Schema:   &JobSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	// http.HandleFunc("/graphql", corsMiddleware(graphQLHandler))
	core.AddEndpoint("/graphql", graphqlHandler)
}

func graphQLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error GraphQLHandler", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error GraphQLHandler", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	params := graphql.Params{Schema: JobSchema, RequestString: string(body)}
	result := graphql.Do(params)
	json.NewEncoder(w).Encode(result)
}
