package server

import (
	"net/http"
	"testing"
)

func TestUnauthorizedHandler(t *testing.T) {
	authHandler := AuthHandler(mockHandler)
	config := ResponseRequestRecorder(http.MethodPost, "http://localhost:8080/api/")
	sctx := ServerContext{W: config.resp, R: config.req, DB: &mockDB}
	authHandler(&sctx)
	if config.resp.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("got %d, want: %d", config.resp.Result().StatusCode, http.StatusUnauthorized)
	}
}

func TestAuthorizedHandler(t *testing.T) {
	authHandler := AuthHandler(mockHandler)
	config := ResponseRequestRecorder(http.MethodPost, "http://localhost:8080/api/")
	config.req.Header.Add("Authorization", "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==")
	sctx := ServerContext{W: config.resp, R: config.req, DB: &mockDB}
	authHandler(&sctx)
	if config.resp.Result().StatusCode != http.StatusOK {
		t.Errorf("got %d, want: %d", config.resp.Result().StatusCode, http.StatusOK)
	}
}
