package db

import (
	"context"
	"goapi/model"
	"log"
	"testing"
)

var (
	ctx context.Context = context.Background()
)

func TestQueryAll(t *testing.T) {
	db, err := StartDatabase()
	rows, err := db.QueryAll(ctx)
	row, err := db.QueryRow(ctx, 1)
	log.Println(rows)
	log.Printf("product: %v \n", row)
	if err != nil {
		t.Errorf("error in database: %v \n", err.Error())
	}
}

func TestConcurrentQueryAll(t *testing.T) {
	ctx := context.Background()
	ch := make(chan []model.Product)
	db, err := StartDatabase()
	go func() {
		rows, _ := db.QueryAll(ctx)
		ch <- rows
	}()
	result := <-ch
	for i := 0; i < len(result); i++ {
		log.Println(result[i])
	}
	if err != nil {
		t.Errorf("error in database")
	}
}