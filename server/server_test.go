package server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	const target string = "http://localhost:8080/api/"

	req := httptest.NewRequest(http.MethodGet, target, nil)
	w := httptest.NewRecorder()

	handler(w, req)

	status := w.Result().Status

	log.Printf("status: %s \n", status)

	if status != "200 OK" {
		t.Errorf("error %s \n", status)
	}
}
