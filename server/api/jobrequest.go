package api

import (
	"jobs/internal/job"
	"jobs/server/config"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type jobRequest struct {
	JobType string `json:"Type" validate:"required,oneof=TIME_CRITICAL NOT_TIME_CRITICAL"`
}

type jobResponse struct {
	ID int `json:"ID"`
}

type errorResponse struct {
	Error string `json:"Error"`
}

type jobStatusResponse struct {
	ID      int    `json:"ID"`
	JobType string `json:"Type"`
	Status  string `json:"Status"`
}

//add job to Queue
func AddJobHandler(dapi *config.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload jobRequest
		var response jobResponse
		if err := preparePayload(r.Body, dapi.Validate, &payload); err != nil {
			log.Printf("AddJobHandler: preparePayload: error: %v, payload: %v", err, payload)
			http.Error(w, "Payload invalid", http.StatusBadRequest)
			return
		}
		id := dapi.Identificator.NextId()
		dapi.Jobs.Enqueue(id, job.JOBSTATUSQUEUED, payload.JobType)
		response.ID = id
		writeResponse(w, response)
	})
}

//get next free job from Queue
func GetJobHandler(dapi *config.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		job := dapi.Jobs.GetNextFreeJob()
		if job == nil {
			var response errorResponse
			response.Error = "Job queue is empty"
			writeResponse(w, response)
			return
		}
		var response jobResponse
		response.ID = job.Item()
		writeResponse(w, response)
	})
}

//Conclude job by id
func ConcludeJobHandler(dapi *config.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		if val, ok := queries["job_id"]; ok {
			strconv.Atoi(val)
			id, _ := strconv.Atoi(val)
			job := dapi.Jobs.ConcludeJobBiId(id)
			if job != nil {
				var response jobResponse
				response.ID = job.Item()
				writeResponse(w, response)
				return
			}
		}

		var response errorResponse
		response.Error = "Job not found"
		writeResponse(w, response)
		return
	})
}

//Get job status by id
func GetJobStatusHandler(dapi *config.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		if val, ok := queries["job_id"]; ok {
			strconv.Atoi(val)
			id, _ := strconv.Atoi(val)
			job := dapi.Jobs.FindBiId(id)
			if job != nil {
				var response jobStatusResponse
				response.ID = job.Item()
				response.Status = job.Status()
				response.JobType = job.JobType()
				writeResponse(w, response)
				return
			}
		}

		var response errorResponse
		response.Error = "Job not found"
		writeResponse(w, response)
		return
	})
}
