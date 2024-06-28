package userclient

import (
	"fmt"

	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	client pb.UserServiceClient
}

func NewUserServiceClient(address string) *UserServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &UserServiceClient{client: pb.NewUserServiceClient(conn)}
}
