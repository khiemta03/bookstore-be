package authclient

import (
	"context"

	httpError "github.com/khiemta03/bookstore-be/api-getway/internal/error"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/authentication"
)

func (c *AuthenticationServiceClient) Login(req *pb.LoginRequest) (*pb.LoginResponse, *httpError.HTTPError) {
	res, err := c.client.Login(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *AuthenticationServiceClient) Register(req *pb.RegisterRequest) (*pb.LoginResponse, *httpError.HTTPError) {
	res, err := c.client.Register(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *AuthenticationServiceClient) ValidateToken(req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, *httpError.HTTPError) {
	res, err := c.client.ValidateToken(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *AuthenticationServiceClient) RenewAccessToken(req *pb.RenewAccessTokenRequest) (*pb.RenewAccessTokenResponse, *httpError.HTTPError) {
	res, err := c.client.RenewAccessToken(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}
