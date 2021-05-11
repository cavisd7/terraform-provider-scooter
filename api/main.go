package main

import (
	"fmt"

	"github.com/cavisd7/terraform-provider-scooter/api/server"
)

func main() {
	items := map[string]server.Item{}

	itemService := server.NewService("localhost:3001", items)
	itemService.Serve()
	if err != nil {
		fmt.Println(err)
	}
}
