package server

import (
	"log"
	"context"
	"goapi/model"
	"net/http"
	"net/http/httptest"
	"testing"
	// "strings"
	"io/ioutil"
)

// MockProducts  used to mock database when calling in routes
type MockProducts struct{}

func (mp *MockProducts) QueryAll(ctx context.Context) ([]model.Product, error) {
	products := []model.Product{
		model.Product{1, "shoes", 100},
		model.Product{2, "pants", 10},
	}
	return products, nil
}

func (mp *MockProducts) QueryRow(ctx context.Context, id int) (model.Product, error) {
	return model.Product{1, "socks", 200}, nil
}

// Init variables for testing
var (
	myServer server       = server{}
	mockDB   MockProducts = MockProducts{}
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

func handler(ctx *serverContext) error {
	_, err := ctx.w.Write([]byte(`hello world`))
	if err != nil {
		return err
	}
	return nil
}

func TestHTTPMethods(t *testing.T) {
	myServer.POST("/api/", handler)
	h := myServer.handlerServer(&mockDB)
	config := ResponseRequestRecorder(http.MethodPost, "http://localhost:8080/api/")
	h(config.resp, config.req)

	if config.resp.Result().StatusCode != http.StatusOK {
		t.Errorf("got %d, want: %d", config.resp.Result().StatusCode, http.StatusOK)
	}
	data, err := ioutil.ReadAll(config.resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if string(data) != "hello world" {
		t.Errorf("got %s, want: %s", data, "hello world")
	}
}

func TestHTTPGetMethodDB(t *testing.T) {
}
