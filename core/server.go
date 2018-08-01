package core

import (
	"log"
	"net/http"
	"strconv"
)

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

func AddEndpoint(path string, handler http.Handler) {
	http.Handle(path, handler)
}

// StartServer Starts graphql api on selected port
func StartServer(port int) {
	stringPort := strconv.Itoa(port)
	err := http.ListenAndServe(":"+stringPort, nil)
	if err != nil {
		log.Println(err)
	}
}
