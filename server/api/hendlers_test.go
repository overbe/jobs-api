package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	type response struct {
		status int
		body   []byte
	}

	tests := []struct {
		name     string
		args     args
		response response
	}{
		{
			name: "http.StatusOK",
			args: args{r: &http.Request{
				Method: "GET",
			},
			},
			response: response{status: http.StatusOK, body: []byte(`{"buildTime":"","gitCommit":"","version":""}` + "\n")},
		},
	}
	for _, tt := range tests {
		req, err := http.NewRequest("GET", "/ping", nil) //nolint: noctx
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(PingHandler)
		handler.ServeHTTP(w, req)

		if gotStatus := w.Result().StatusCode; gotStatus != tt.response.status { //nolint: bodyclose
			t.Errorf("handler returned wrong status code: got %v want %v",
				gotStatus, tt.response.status)
		}
		if !bytes.Equal(w.Body.Bytes(), tt.response.body) {
			t.Errorf("handler returned unexpected body: got %s want %s", w.Body.Bytes(), tt.response.body)
		}
	}
}
