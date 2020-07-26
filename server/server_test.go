package server

import (
	"log"
	"context"
	"goapi/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
)

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

func TestGet(t *testing.T) {
	const target string = "http://localhost:8080/api/"
	mockModel := MockProducts{}
	req := httptest.NewRequest(http.MethodGet, target, nil)
	w := httptest.NewRecorder()
	handler(w, req, &mockModel)
	result := w.Result()
	status := result.Status
	body := result.Body
	data, err := ioutil.ReadAll(body)
	log.Printf("%s\n", data)
	if status != "200 OK" {
		t.Errorf("error %s \n", status)
	}

	if err != nil {
		t.Errorf("error %s \n", err.Error())
	}
}
