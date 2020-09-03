package server

import (
	"context"
	"jobs/server/api"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"jobs/server/config"
	"log"
	"net/http"
	"strings"
)

//Run server
func Run(s *config.Builder) error {
	var r = mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/ping", api.PingHandler).Methods("GET")

	apiEndpoints := r.PathPrefix("/api/v1").Subrouter()
	apiEndpoints.Handle("/jobs/enqueue", api.AddJobHandler(s)).Methods("POST")
	apiEndpoints.Handle("/jobs/dequeue", api.GetJobHandler(s)).Methods("POST")
	apiEndpoints.Handle("/jobs/{job_id}/conclude", api.ConcludeJobHandler(s)).Methods("POST")
	apiEndpoints.Handle("/jobs/{job_id}", api.GetJobStatusHandler(s)).Methods("GET")

	err := r.Walk(walk)
	if err != nil {
		log.Println(err)
	}
	return ListenAndServe(s, r)
}

// ListenAndServe start web framework.
func ListenAndServe(s *config.Builder, h http.Handler) error {
	srv := &http.Server{
		Handler:      h,
		Addr:         "0.0.0.0:" + s.Cfg.Server.Port,
		WriteTimeout: time.Duration(s.Cfg.Server.WriteTimeoutSec) * time.Second,
		ReadTimeout:  time.Duration(s.Cfg.Server.ReadTimeoutSec) * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("server initialization error: %v", err)
		}
	}()
	log.Println("JOBS server Started")

	<-stop

	log.Printf("Waiting at most %v seconds for a graceful shutdown", time.Second*time.Duration(s.Cfg.Server.ShutdownTime))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(s.Cfg.Server.ShutdownTime))
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("JOBS server finalization error: %v", err)
	}
	log.Println("JOBS server exited Properly")
	return nil
}

//Check all rotutes and methods
func walk(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	pathTemplate, err := route.GetPathTemplate()
	if err == nil {
		log.Println("ROUTE:", pathTemplate)
	}
	pathRegexp, err := route.GetPathRegexp()
	if err == nil {
		log.Println("Path regexp:", pathRegexp)
	}
	methods, err := route.GetMethods()
	if err == nil {
		log.Println("Methods:", strings.Join(methods, ","))
	}
	log.Println()
	return nil
}
