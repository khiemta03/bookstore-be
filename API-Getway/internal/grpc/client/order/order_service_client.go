package orderclient

import (
	"fmt"

	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderServiceClient struct {
	client pb.OrderServiceClient
}

func NewOrderServiceClient(address string) *OrderServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &OrderServiceClient{client: pb.NewOrderServiceClient(conn)}
}
