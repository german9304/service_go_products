package db

import (
	"log"
	"testing"
)

func TestDB(t *testing.T) {

	db, err := Start()
	rows, err := db.QueryAll()

	p := Product{}

	log.Println(p.Products(rows))

	if err != nil {
		t.Errorf("error in database")
	}
}
