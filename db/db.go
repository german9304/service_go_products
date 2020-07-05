package db

import (
	"database/sql"
	// "log"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func (sqlDB *DB) QueryAll() (*sql.Rows, error) {
	rows, err := sqlDB.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Start() (*DB, error) {
	connStr := "user=postgres dbname=mydb password='user123456' sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}
