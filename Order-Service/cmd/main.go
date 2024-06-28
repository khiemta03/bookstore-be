package main

import (
	"log"

	"github.com/khiemta03/bookstore-be/order-service/api"
	"github.com/khiemta03/bookstore-be/order-service/internal/config"
)

func main() {
	config, err := config.LoadConfig("./configs", "dev", "env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server, err := api.NewServerWithConfig(&config)
	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}
