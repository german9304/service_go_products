package server

import (
	"net/http"
	"net/http/httptest"
)

var (
	mockDB MockProducts = MockProducts{}
)

type RequestConfig struct {
	req  *http.Request
	resp *httptest.ResponseRecorder
}

func ResponseRequestRecorder(method, url string) RequestConfig {
	req := httptest.NewRequest(method, url, nil)
	resp := httptest.NewRecorder()
	return RequestConfig{req, resp}
}

func mockHandler(ctx *ServerContext) error {
	_, err := ctx.W.Write([]byte(`hello world`))
	if err != nil {
		return err
	}
	return nil
}
