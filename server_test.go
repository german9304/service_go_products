package server

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	// "strings"
	"io/ioutil"
)

// Init variables for testing
var (
	myServer Server = Server{}
)

type ProductsResponse struct {
	Products []Product `json:"products"`
}

func databaseHanlder(ctx *ServerContext) error {
	rows, err := ctx.DB.QueryAll()
	if err != nil {
		return err
	}
	return ctx.JSON(ProductsResponse{rows})
}

func TestHTTPMethods(t *testing.T) {
	myServer.POST("/api/", mockHandler)
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
	if err != nil {
		log.Fatal(err)
	}
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
