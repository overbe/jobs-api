package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"jobs/internal/platform/validate"
	"log"
	"net/http"
)

//nolint:gochecknoglobals
var version, gitCommit, buildTime string

func PingHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(map[string]string{
		"version":   version,
		"buildTime": buildTime,
		"gitCommit": gitCommit,
	})
	if err != nil {
		log.Printf("Ping with error %v", err)
	}
}

func preparePayload(r io.Reader, v *validate.Validator, payload interface{}) error {
	rawPayload, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(rawPayload, &payload); err != nil {
		return err
	}
	if err := v.Validate.Struct(payload); err != nil {
		return err
	}
	return nil
}

func writeResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("writeResponse with error %v", err)
	}
}
