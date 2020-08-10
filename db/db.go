package db

import (
	"context"
	"goapi/model"
	"log"

	pgx "github.com/jackc/pgx/v4"
	"github.com/rs/xid"
)

// IDB Defines the Product interface, that receives a context
type IDB interface {
	QueryAll() ([]model.Product, error)
	QueryRow(id string) (model.Product, error)
	CreateRow(name string, price int) (model.Product, error)
	DeleteRow(id string) (string, error)
}

type DB struct {
	db  *pgx.Conn
	ctx context.Context
}

// QueryAll elements from db
func (sqlDB *DB) QueryAll() ([]model.Product, error) {
	const SQL_STATEMENT = `
	SELECT id, name, price 
	FROM products
	`
	rows, err := sqlDB.db.Query(sqlDB.ctx, SQL_STATEMENT)
	if err != nil {
		return nil, err
	}
	select {
	case <-sqlDB.ctx.Done():
		log.Println("there is an error in the database")
		return nil, sqlDB.ctx.Err()
	default:
		log.Println("no error in the database")
	}
	return model.Products(rows)
}

// Queries a single row
func (sqlDB *DB) QueryRow(id string) (model.Product, error) {
	const SQL_STATEMENT = `
	SELECT id, name, price 
	FROM Products WHERE id = $1
	`
	row := sqlDB.db.QueryRow(sqlDB.ctx, SQL_STATEMENT, id)
	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// CreateRow creates a new row in the database
func (sqlDB *DB) CreateRow(name string, price int) (model.Product, error) {
	const SQL_STATEMENT = `
	INSERT INTO products (id, name, price)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	guid := xid.New()
	var productID string
	row := sqlDB.db.QueryRow(sqlDB.ctx, SQL_STATEMENT, guid, name, price)
	err := row.Scan(&productID)
	if err != nil {
		return model.Product{}, err
	}
	return model.New(guid.String(), name, price), nil
}

// DeleteRow deletes a row in the database
func (sqlDB *DB) DeleteRow(id string) (string, error) {
	const SQL_STATEMENT = `
	DELETE FROM products 
	WHERE id = $1
	RETURNING id
	`
	var productID string
	row := sqlDB.db.QueryRow(sqlDB.ctx, SQL_STATEMENT, id)
	err := row.Scan(&productID)
	if err != nil {
		return "", err
	}
	return id, nil
}

func StartDatabase(ctx context.Context) (DB, error) {
	const URL = "postgres://user@test:testing123@database:5432/mydb?sslmode=disable"
	db, err := pgx.Connect(ctx, URL)
	if err != nil {
		return DB{}, err
	}
	return DB{db, ctx}, nil
}
