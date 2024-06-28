package client

import (
	"fmt"

	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func NewUserServiceClient(address string) UserServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return UserServiceClient{Client: pb.NewUserServiceClient(conn)}
}
