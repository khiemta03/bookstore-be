package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/khiemta03/bookstore-be/user-service/internal/config"
	db "github.com/khiemta03/bookstore-be/user-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/gapi"
)

func main() {
	config, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	store := db.NewStore(conn)
	server := gapi.NewServer(store)

	address := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	server.Run(address)
}
