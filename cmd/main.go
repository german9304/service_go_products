package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	service "github.com/goapi"

	"github.com/joho/godotenv"
)

func productsHandler(ctx *service.ServerContext) error {
	rows, err := ctx.DB.QueryAll()
	type ProductsResponse struct {
		Products []service.Product `json:"products"`
	}
	if err != nil {
		return err
	}
	return ctx.JSON(ProductsResponse{rows})
}

func productHandler(ctx *service.ServerContext) error {
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

func createProductHandler(ctx *service.ServerContext) error {
	type Data struct {
		Name  string
		Price int
	}
	var data Data
	ctx.R.Header.Set("Content-type", "application/json")
	body := ctx.R.Body
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	newProduct, err := ctx.DB.CreateRow(data.Name, data.Price)
	if err != nil {
		return err
	}
	return ctx.JSON(newProduct)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	myServer := service.Server{}
	myServer.GET("/api/products", productsHandler)
	myServer.GET("/api/product", productHandler)
	myServer.POST("/api/product", createProductHandler)
	err = myServer.Run("8080")
	if err != nil {
		log.Println(err)
	}
}
