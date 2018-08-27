package rest

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/interfaces/rest/jobs"
)

func init() {
	router := getRouterInstance()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	core.AddEndpoint("/", router)
}

func getRouterInstance() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/jobs", jobs.Routes())
	})

	return router
}
