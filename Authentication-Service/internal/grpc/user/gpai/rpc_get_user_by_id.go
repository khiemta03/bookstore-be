package client

import (
	"context"
	"fmt"

	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/user/pb"
	errors "github.com/khiemta03/bookstore-be/authentication-service/pkg/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *UserServiceClient) GetUserById(req *pb.GetUserByIdRequest) (*pb.User, error) {
	res, err := c.Client.GetUserById(context.Background(), &pb.GetUserByIdRequest{
		Id: req.GetId(),
	})

	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf(errors.NonGRPCError)
		}

		switch st.Code() {
		case codes.NotFound:
			return nil, fmt.Errorf(errors.NotFoundError)
		case codes.Internal:
			return nil, fmt.Errorf(errors.InternalServerError)
		default:
			return nil, fmt.Errorf(errors.UnknownGRPCError)
		}
	}

	return res, nil
}
