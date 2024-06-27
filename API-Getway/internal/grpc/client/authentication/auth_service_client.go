package authclient

import (
	"fmt"

	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthenticationServiceClient struct {
	client pb.AuthenticationServiceClient
}

func NewAuthenticationServiceClient(address string) *AuthenticationServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &AuthenticationServiceClient{client: pb.NewAuthenticationServiceClient(conn)}
}
