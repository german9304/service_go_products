package auth

import (
	"context"
	//	"encoding/json"
	"goapi/mock"
	"goapi/server"
	//	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	// "strings"
	//	"io/ioutil"
)

var (
	mockDB mock.MockProducts = mock.MockProducts{}
	ctx    context.Context   = context.Background()
)

func handler(ctx *server.ServerContext) error {
	_, err := ctx.W.Write([]byte(`hello world`))
	if err != nil {
		return err
	}
	return nil
}

type RequestConfig struct {
	req  *http.Request
	resp *httptest.ResponseRecorder
}

func ResponseRequestRecorder(method, url string) RequestConfig {
	req := httptest.NewRequest(method, url, nil)
	resp := httptest.NewRecorder()
	return RequestConfig{req, resp}
}

func TestHandler(t *testing.T) {
	authHandler := Handler(handler)
	config := ResponseRequestRecorder(http.MethodPost, "http://localhost:8080/api/")
	sctx := server.ServerContext{config.resp, config.req, &mockDB, ctx}
	authHandler(&sctx)
	if config.resp.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("got %d, want: %d", config.resp.Result().StatusCode, http.StatusUnauthorized)
	}
}

func TestAuthHandler(t *testing.T) {
	authHandler := Handler(handler)
	config := ResponseRequestRecorder(http.MethodPost, "http://localhost:8080/api/")
	config.req.Header.Add("Authorization", "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==")
	sctx := server.ServerContext{config.resp, config.req, &mockDB, ctx}
	authHandler(&sctx)
	if config.resp.Result().StatusCode != http.StatusOK {
		t.Errorf("got %d, want: %d", config.resp.Result().StatusCode, http.StatusOK)
	}
}
