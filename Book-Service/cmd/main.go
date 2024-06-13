package main

import (
	"fmt"
	"log"

	"github.com/khiemta03/bookstore-be/book-service/api"
	"github.com/khiemta03/bookstore-be/book-service/internal/config"
	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
)

func main() {
	config, err := config.LoadConfig("./configs")
	if err != nil {
		// TODO: send err to log service
		log.Fatal("cannot load config:", err)
	}

	store, err := db.NewStore(config.DBDriver, config.DBSource)
	if err != nil {
		// TODO: send err to log service
		log.Fatal("cannot connect to the database:", err)
	}

	server := api.NewServer(store)

	address := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	server.Run(address)
}
