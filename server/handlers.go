package server

import (
	"encoding/json"
	"goapi/model"
	"net/http"

	mydb "goapi/db"
)

func getProducts(w http.ResponseWriter, _ *http.Request, db mydb.IDB) error {
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
