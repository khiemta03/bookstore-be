package bookclient

import (
	"fmt"

	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookServiceClient struct {
	client pb.BookServiceClient
}

func NewBookServiceClient(address string) *BookServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &BookServiceClient{client: pb.NewBookServiceClient(conn)}
}
