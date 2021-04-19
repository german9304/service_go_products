package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goapi/mock"
	"github.com/goapi/model"

	// "strings"
	"io/ioutil"
)

// Init variables for testing
var (
	myServer Server            = Server{}
	mockDB   mock.MockProducts = mock.MockProducts{}
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
	myServer.GET("/api/", databaseHanlder)
	h := myServer.handlerServer(&mockDB)
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

func TestMethodNotAllowed(t *testing.T) {
	myServer.GET("/api/", databaseHanlder)
	h := myServer.handlerServer(&mockDB)
	config := ResponseRequestRecorder(http.MethodPut, "http://localhost:8080/api/")
	h(config.resp, config.req)
	log.Printf("response status: %d \n", config.resp.Result().StatusCode)
	responseStatusCode := config.resp.Result().StatusCode
	if responseStatusCode != 405 {
		t.Errorf("got: %d want %d \n", responseStatusCode, 405)
	}
}
