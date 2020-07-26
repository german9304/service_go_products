package db

import (
	"context"
	"database/sql"
	"goapi/model"
	"log"

	_ "github.com/lib/pq"
)

// IDB Defines the Product interface, that receives a context
type IDB interface {
	QueryAll(ctx context.Context) ([]model.Product, error)
	QueryRow(ctx context.Context, id int) (model.Product, error)
}

type DB struct {
	db *sql.DB
}

// QueryAll elements from db
func (sqlDB *DB) QueryAll(ctx context.Context) ([]model.Product, error) {
	rows, err := sqlDB.db.QueryContext(ctx, "SELECT id, name, price FROM products")
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

func (sqlDB *DB) QueryRow(ctx context.Context, id int) (model.Product, error) {
	row := sqlDB.db.QueryRowContext(ctx, "SELECT id, name, price FROM Products WHERE id = $1", id)
	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func StartDatabase() (DB, error) {
	connStr := "user=postgres dbname=mydb password='user123456' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return DB{}, err
	}
	return DB{db}, nil
}
