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
		id := dapi.Identificator.NextID()
		dapi.Jobs.Enqueue(id, job.JOBSTATUSQUEUED, payload.JobType)
		response.ID = id
		writeResponse(w, response)
	})
}

//get next free job from Queue
func GetJobHandler(dapi *config.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		freeJob := dapi.Jobs.GetNextFreeJob()
		if freeJob == nil {
			var response errorResponse
			response.Error = "Job queue is empty"
			writeResponse(w, response)
			return
		}
		var response jobResponse
		response.ID = freeJob.Item()
		writeResponse(w, response)
	})
}

//Conclude job by id
func ConcludeJobHandler(dapi *config.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		if val, ok := queries["job_id"]; ok {
			id, _ := strconv.Atoi(val)
			concludeJob := dapi.Jobs.ConcludeJobBiID(id)
			if concludeJob != nil {
				var response jobResponse
				response.ID = concludeJob.Item()
				writeResponse(w, response)
				return
			}
		}

		var response errorResponse
		response.Error = "Job not found"
		writeResponse(w, response)
	})
}

//Get job status by id
func GetJobStatusHandler(dapi *config.Builder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		if val, ok := queries["job_id"]; ok {
			id, err := strconv.Atoi(val)
			if err != nil {
				log.Printf("GetJobStatusHandler: preparePayload: error: %v", err)
				http.Error(w, "Payload invalid", http.StatusBadRequest)
				return
			}
			foundJob := dapi.Jobs.FindByID(id)
			if foundJob != nil {
				var response jobStatusResponse
				response.ID = foundJob.Item()
				response.Status = foundJob.Status()
				response.JobType = foundJob.JobType()
				writeResponse(w, response)
				return
			}
		}

		var response errorResponse
		response.Error = "Job not found"
		writeResponse(w, response)
	})
}
