package main

import (
	"log"
	"goapi/server"
)

func main() {
	err := server.Run("8080")
	if err != nil {
		log.Println(err)
	}
}
