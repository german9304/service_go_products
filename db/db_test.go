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
	db, err := Start(ctx)
	if err != nil {
		t.Errorf("error in database: %v \n", err.Error())
	}
	_, errQAll := db.QueryAll()
	if errQAll != nil {
		t.Errorf("error in database: %v \n", errQAll.Error())
	}
}

func TestConcurrentQueryAll(t *testing.T) {
	ctx := context.Background()
	ch := make(chan []model.Product)
	db, err := Start(ctx)
	go func() {
		rows, _ := db.QueryAll()
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

func TestCreateRow(t *testing.T) {
	ctx := context.Background()
	db, err := Start(ctx)
	name := "furniture"
	price := 430
	newProduct, err := db.CreateRow(name, price)
	if newProduct.Name != name {
		t.Errorf("got: %s, want: %s \n", newProduct.Name, name)
	}
	_, err = db.DeleteRow(newProduct.ID)
	if err != nil {
		t.Errorf("error: %s ", err)
	}
}
