package db

import (
	"database/sql"
)

// Product interface
type Product struct {
	ID    int
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func New(name string, id, price int) Product {
	return Product{
		id,
		name,
		price,
	}
}

func (p *Product) Products(rows *sql.Rows) ([]Product, error) {
	var Products []Product

	for rows.Next() {
		var newProduct Product = Product{}
		if err := rows.Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price); err != nil {
			return nil, err
		}
		Products = append(Products, newProduct)
	}

	return Products, nil
}
