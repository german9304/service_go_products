package model

import "database/sql"

// Product data
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func New(id, name string, price int) Product {
	return Product{
		id,
		name,
		price,
	}
}

func Products(rows *sql.Rows) ([]Product, error) {
	var products []Product

	for rows.Next() {
		var newProduct Product = Product{}
		if err := rows.Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price); err != nil {
			return nil, err
		}
		products = append(products, newProduct)
	}

	return products, nil
}
