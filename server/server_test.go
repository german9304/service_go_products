package server

import (
	"context"
	"encoding/json"
	"goapi/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	// "strings"
	"io/ioutil"
)

// MockProducts  used to mock database when calling in routes
type MockProducts struct{}

func (mp *MockProducts) QueryAll() ([]model.Product, error) {
	products := []model.Product{
		model.Product{"3o3o3", "shoes", 100},
		model.Product{"3020222", "pants", 10},
	}
	return products, nil
}

func (mp *MockProducts) QueryRow(id string) (model.Product, error) {
	return model.Product{"123333", "socks", 200}, nil
}

// Init variables for testing
var (
	myServer Server       = Server{}
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

func handler(ctx *ServerContext) error {
	_, err := ctx.W.Write([]byte(`hello world`))
	if err != nil {
		return err
	}
	return nil
}

type ProductsResponse struct {
	Products []model.Product `json:"products"`
}

func databaseHanlder(ctx *ServerContext) error {
	rows, err := ctx.DB.QueryAll()
	if err != nil {
		return err
	}
	return ctx.JSON(ProductsResponse{rows})
}

func TestHTTPMethods(t *testing.T) {
	ctx := context.Background()
	myServer.POST("/api/", handler)
	h := myServer.handlerServer(&mockDB, ctx)
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
	ctx := context.Background()
	myServer.GET("/api/", databaseHanlder)
	h := myServer.handlerServer(&mockDB, ctx)
	config := ResponseRequestRecorder(http.MethodGet, "http://localhost:8080/api/")
	h(config.resp, config.req)
	bodyData, err := ioutil.ReadAll(config.resp.Body)
	var data ProductsResponse
	err = json.Unmarshal(bodyData, &data)
	if err != nil {
		log.Fatal(err)
	}
	if len(data.Products) != 2 {
		t.Errorf("got: %d, want: %d ", len(data.Products), 2)
	}
	content := config.resp.Result().Header.Get("Content-type")
	if content != "application/json" {
		t.Errorf("got: %s, want: %s ", content, "application/json")
	}
}
