package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/plopezm/kaiser/core/provider/interfaces"

	"github.com/go-chi/render"
	"github.com/plopezm/kaiser/core"
	"github.com/plopezm/kaiser/core/types"

	"github.com/go-chi/chi"
)

// Routes Returns all the package routes
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetJobs)
	router.Post("/", CreateJob)
	router.Post("/{jobName}", ExecuteJob)
	return router
}

// GetJobs Returns all stored jobs
func GetJobs(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, core.GetEngineInstance().GetJobs())
}

// CreateJob Creates a new job
func CreateJob(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var job types.Job
	err := decoder.Decode(&job)

	response := make(map[string]interface{})
	if err != nil {
		response["error"] = err.Error()
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response)
		return
	}

	err = interfaces.NotifyJob(&job)
	if err != nil {
		response["error"] = err.Error()
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response)
		return
	}

	response["message"] = "Job creation success"
	response["job"] = job
	render.JSON(w, r, response)
}

// ExecuteJob Executes an existing job
func ExecuteJob(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	jobName := chi.URLParam(r, "jobName")
	decoder := json.NewDecoder(r.Body)
	var params map[string]interface{}
	err := decoder.Decode(&params)
	if err != nil {
		response["error"] = err.Error()
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, response)
		return
	}

	engineInstance := core.GetEngineInstance()
	err = engineInstance.ExecuteStoredJob(jobName, params)

	if err != nil {
		response["error"] = err.Error()
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, response)
		return
	}

	response["message"] = "Job [" + jobName + "] Executed"
	render.JSON(w, r, response)
}
