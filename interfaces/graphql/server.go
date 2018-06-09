package graphql

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
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

// StartServer Starts graphql api on selected port
func StartServer(port int) {
	stringPort := strconv.Itoa(port)
	http.HandleFunc("/api/graphql", graphQLHandler)
	log.Println("Started GraphQL interface on " + stringPort)
	err := http.ListenAndServe(":"+stringPort, nil)
	if err != nil {
		log.Println(err)
	}
}
