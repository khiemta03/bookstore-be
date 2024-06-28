package main

import (
	"log"

	"github.com/khiemta03/bookstore-be/api-getway/api"
	"github.com/khiemta03/bookstore-be/api-getway/internal/config"
)

func main() {
	config, err := config.LoadConfig("./configs", "dev", "env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := api.NewServer(&config)

	server.InitRoutes()

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
