package client

import (
	"context"

	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/user/pb"
)

func (c *UserServiceClient) CreateUser(req *pb.CreateUserRequest) (*pb.User, error) {
	res, err := c.Client.CreateUser(context.Background(), &pb.CreateUserRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
		Fullname: req.GetEmail(),
		Phone:    req.GetPhone(),
		Address:  req.GetAddress(),
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
