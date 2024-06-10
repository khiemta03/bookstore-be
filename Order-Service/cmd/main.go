package main

import (
	"fmt"
	"log"

	"github.com/khiemta03/bookstore-be/order-service/api"
	"github.com/khiemta03/bookstore-be/order-service/internal/config"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
)

func main() {
	// conn, err := grpc.NewClient("localhost:3003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Println("Could not connect:", err)
	// }

	// client := pb.NewBookServiceClient(conn)

	// result, err := client.GetBook(context.Background(), &pb.GetBookRequest{
	// 	// Id: "5a96b419-5902-48f0-934e-57b066c8f69c",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result.Id, result.Title, result.FullTitle, result.Price, result.StockQuantity)
	// 	fmt.Println(result)
	// }
	config, err := config.LoadConfig("./configs", "dev", "env")
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
