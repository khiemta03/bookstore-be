package client

import (
	"context"

	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/user/pb"
)

func (c *UserServiceClient) GetUserByName(req *pb.GetUserByNameRequest) (*pb.User, error) {
	res, err := c.Client.GetUserByName(context.Background(), &pb.GetUserByNameRequest{
		Username: req.GetUsername(),
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
