package main

import (
	"log"

	"github.com/cavisd7/terraform-provider-scooter/api/server"
)

func main() {
	items := map[string]server.Item{}

	itemService := server.NewService("localhost:3001", items)
	err := itemService.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
