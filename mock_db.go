package server

// MockProducts  used to mock database when calling in routes
type MockProducts struct{}

func (mp *MockProducts) QueryAll() ([]Product, error) {
	products := []Product{
		{"3o3o3", "shoes", 100},
		{"3020222", "pants", 10},
	}
	return products, nil
}

func (mp *MockProducts) QueryRow(id string) (Product, error) {
	return Product{"123333", "socks", 200}, nil
}

func (mp *MockProducts) CreateRow(name string, price int) (Product, error) {
	return Product{"1223", name, price}, nil
}

func (mp *MockProducts) DeleteRow(id string) (string, error) {
	return "123", nil
}
