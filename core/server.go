package core

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}).Handler,
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
}

// AddEndpoint Adds a http endpoint to the server
func AddEndpoint(path string, handler http.Handler) {
	router.Mount(path, handler)
}

// StartServer Starts graphql api on selected port
func StartServer(port int) {
	stringPort := strconv.Itoa(port)
	log.Fatal(http.ListenAndServe(":"+stringPort, router))
}
