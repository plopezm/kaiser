package rest

import (
	"github.com/go-chi/chi"
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/interfaces/rest/jobs"
)

func init() {
	router := getRouterInstance()
	core.AddEndpoint("/", router)
}

func getRouterInstance() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/jobs", jobs.Routes())
	})

	return router
}
