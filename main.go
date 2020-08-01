package main

import (
	"goapi/model"
	"log"
	"goapi/server"
)

func main() {
	myServer := server.Server{}
	myServer.GET("/api/products", func(ctx *server.ServerContext) error {
		rows, err := ctx.DB.QueryAll()
		type ProductsResponse struct {
			Products []model.Product `json:"products"`
		}		
		if err != nil {
			return err
		}
		return ctx.JSON(ProductsResponse{rows})
	})

	err := myServer.Run("8080")
	if err != nil {
		log.Println(err)
	}
}
