package main

import (
	"encoding/json"
	"goapi/model"
	"goapi/server"
	"io/ioutil"
	"log"
)

type Message struct {
	Message string `json:"message"`
}

func productsHandler(ctx *server.ServerContext) error {
	rows, err := ctx.DB.QueryAll()
	type ProductsResponse struct {
		Products []model.Product `json:"products"`
	}
	if err != nil {
		return err
	}
	return ctx.JSON(ProductsResponse{rows})
}

func productHandler(ctx *server.ServerContext) error {
	query := ctx.R.URL.Query()
	_, ok := query["id"]
	if ok {
		product, err := ctx.DB.QueryRow(query.Get("id"))
		if err != nil {
			return err
		}
		return ctx.JSON(product)
	}
	return nil
}

func createProductHandler(ctx *server.ServerContext) error {
	type Data struct {
		Name  string
		Price int
	}
	var data Data
	ctx.R.Header.Set("Content-type", "application/json")
	body := ctx.R.Body
	b, err := ioutil.ReadAll(body)
	err = json.Unmarshal(b, &data)
	newProduct, err := ctx.DB.CreateRow(data.Name, data.Price)
	if err != nil {
		return err
	}
	return ctx.JSON(newProduct)
}

func main() {
	myServer := server.Server{}
	myServer.GET("/api/products", productsHandler)
	myServer.GET("/api/product", productHandler)
	myServer.POST("/api/product", createProductHandler)
	err := myServer.Run("8080")
	if err != nil {
		log.Println(err)
	}
}
