package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goapi/server"

	"github.com/goapi/mock"
)

var (
	mockDB mock.MockProducts = mock.MockProducts{}
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

func TestUnauthorizedHandler(t *testing.T) {
	authHandler := Handler(handler)
	config := ResponseRequestRecorder(http.MethodPost, "http://localhost:8080/api/")
	sctx := server.ServerContext{W: config.resp, R: config.req, DB: &mockDB}
	authHandler(&sctx)
	if config.resp.Result().StatusCode != http.StatusUnauthorized {
		t.Errorf("got %d, want: %d", config.resp.Result().StatusCode, http.StatusUnauthorized)
	}
}

func TestAuthorizedHandler(t *testing.T) {
	authHandler := Handler(handler)
	config := ResponseRequestRecorder(http.MethodPost, "http://localhost:8080/api/")
	config.req.Header.Add("Authorization", "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==")
	sctx := server.ServerContext{W: config.resp, R: config.req, DB: &mockDB}
	authHandler(&sctx)
	if config.resp.Result().StatusCode != http.StatusOK {
		t.Errorf("got %d, want: %d", config.resp.Result().StatusCode, http.StatusOK)
	}
}
