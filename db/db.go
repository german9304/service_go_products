package db

import (
	"context"
	"goapi/model"
	"log"

	pgx "github.com/jackc/pgx/v4"
)

// IDB Defines the Product interface, that receives a context
type IDB interface {
	QueryAll(ctx context.Context) ([]model.Product, error)
	QueryRow(ctx context.Context, id int) (model.Product, error)
}

type DB struct {
	db *pgx.Conn
}

// QueryAll elements from db
func (sqlDB *DB) QueryAll(ctx context.Context) ([]model.Product, error) {
	rows, err := sqlDB.db.Query(ctx, "SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		log.Println("there is an error in the database")
		return nil, ctx.Err()
	default:
		log.Println("no error in the database")
	}
	return model.Products(rows)
}

// Queries a single row 
func (sqlDB *DB) QueryRow(ctx context.Context, id int) (model.Product, error) {
	row := sqlDB.db.QueryRow(ctx, "SELECT id, name, price FROM Products WHERE id = $1", id)
	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func StartDatabase(ctx context.Context) (DB, error) {
	const URL = "postgres://postgres:user123456@localhost:5432/mydb?sslmode=disable"
	db, err := pgx.Connect(ctx, URL)
	if err != nil {
		return DB{}, err
	}
	return DB{db}, nil
}
