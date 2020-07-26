package server

import (
	"context"
	"encoding/json"
	mydb "goapi/db"
	"goapi/model"
	//"io/ioutil"
	"log"
	"net/http"
)

var (
	ctx context.Context = context.Background()
)

type handlerContext func(w http.ResponseWriter, r *http.Request, db mydb.IDB) error

func handler(w http.ResponseWriter, _ *http.Request, db mydb.IDB) error {
	w.Header().Set("Content-type", "application/json")
	rows, err := db.QueryAll(ctx)
	if err != nil {
		return err
	}
	type ProductsResponse struct {
		Products []model.Product `json:"products"`
	}
	b, err := json.Marshal(ProductsResponse{rows})
	w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func handlerAdapter(h handlerContext, db mydb.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r, &db)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Run(port string) error {
	url := ":" + port
	db, err := mydb.StartDatabase()
	mux := http.NewServeMux()
	mux.Handle("/api/products", handlerAdapter(handler, db))
	log.Printf("listening on port http://localhost%s \n", url)
	err = http.ListenAndServe(url, mux)
	if err != nil {
		return err
	}
	return nil
}
