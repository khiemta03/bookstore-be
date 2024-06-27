package userclient

import (
	"context"
	"io"

	httpError "github.com/khiemta03/bookstore-be/api-getway/internal/error"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/user"
)

func (c *UserServiceClient) GetUserById(req *pb.GetUserByIdRequest) (*pb.User, *httpError.HTTPError) {
	res, err := c.client.GetUserById(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *UserServiceClient) ListUsers(req *pb.ListUsersRequest) ([]*pb.User, *httpError.HTTPError) {
	var responses []*pb.User
	stream, err := c.client.ListUsers(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, httpError.MapGRPCErrorToHTTPError(err)
		}
		responses = append(responses, res)
	}

	return responses, nil
}
