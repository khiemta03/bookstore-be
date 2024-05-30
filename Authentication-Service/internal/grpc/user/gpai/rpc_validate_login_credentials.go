package client

import (
	"context"

	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/user/pb"
)

func (c *UserServiceClient) ValidateLoginCredentials(req *pb.ValidateLoginCredentialsRequest) (*pb.ValidateLoginCredentialsResponse, error) {
	res, err := c.Client.ValidateLoginCredentials(context.Background(), &pb.ValidateLoginCredentialsRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
