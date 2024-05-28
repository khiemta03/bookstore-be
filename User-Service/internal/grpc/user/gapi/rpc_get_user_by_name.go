package gapi

import (
	"context"
	"database/sql"

	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.User, error) {
	username := req.GetUsername()

	user, err := server.store.GetUserByName(ctx, username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, UserNotFoundError)
		}

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	res := convertUser(user)

	return res, nil
}
