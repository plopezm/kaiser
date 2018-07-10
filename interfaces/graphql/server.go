package graphql

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

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

func corsMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		next(w, r)
	}
}

// StartServer Starts graphql api on selected port
func StartServer(port int) {
	stringPort := strconv.Itoa(port)
	graphqlHandler := handler.New(&handler.Config{
		Schema:   &JobSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	// http.HandleFunc("/graphql", corsMiddleware(graphQLHandler))
	http.Handle("/graphql", graphqlHandler)

	log.Println("Started GraphQL interface on " + stringPort)
	err := http.ListenAndServe(":"+stringPort, nil)
	if err != nil {
		log.Println(err)
	}
}
