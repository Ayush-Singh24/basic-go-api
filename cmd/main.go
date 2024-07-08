package main

import (
	"log"

	"github.com/Ayush-Singh24/basic-go-api/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
