package mock

import (
	"github.com/goapi/model"
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

func (mp *MockProducts) CreateRow(name string, price int) (model.Product, error) {
	return model.New("1223", name, price), nil
}

func (mp *MockProducts) DeleteRow(id string) (string, error) {
	return "123", nil
}
