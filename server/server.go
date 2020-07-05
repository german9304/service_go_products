package server

import (
	"encoding/json"
	mydb "goapi/db"
	//"io/ioutil"
	"log"
	"net/http"
)

type handlerContext func(w http.ResponseWriter, r *http.Request, db *mydb.DB) error

func handler(w http.ResponseWriter, _ *http.Request, db *mydb.DB) error {
	w.Header().Set("Content-type", "application/json")

	rows, err := db.QueryAll()
	if err != nil {
		return err
	}
	product := mydb.Product{}
	products, err := product.Products(rows)

	if err != nil {
		return err
	}

	type ProductsResponse struct {
		Products []mydb.Product `json:"products"`
	}

	b, err := json.Marshal(ProductsResponse{products})

	w.Write(b)

	if err != nil {
		return err
	}

	return nil
}

func handlerAdapter(h handlerContext, text string) http.HandlerFunc {
	db, err := mydb.Start()

	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			log.Fatal(err)
		}
		err := h(w, r, db)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Run(port string) {

	mux := http.NewServeMux()

	mux.Handle("/api/products", handlerAdapter(handler, "text"))

	http.ListenAndServe(":"+port, mux)
}
