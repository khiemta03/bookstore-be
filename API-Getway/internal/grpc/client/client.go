package client

import (
	authclient "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/client/authentication"
	bookclient "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/client/book"
	userclient "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/client/user"
)

type Client struct {
	*authclient.AuthenticationServiceClient
	*userclient.UserServiceClient
	*bookclient.BookServiceClient
}

type Config struct {
	UserServiceAddress string `json:"user_service_address"`
	AuthServiceAddress string `json:"auth_service_address"`
	BookServiceAddress string `json:"book_service_address"`
}

func NewClient(config Config) *Client {
	return &Client{
		AuthenticationServiceClient: authclient.NewAuthenticationServiceClient(config.AuthServiceAddress),
		UserServiceClient:           userclient.NewUserServiceClient(config.UserServiceAddress),
		BookServiceClient:           bookclient.NewBookServiceClient(config.BookServiceAddress),
	}
}
