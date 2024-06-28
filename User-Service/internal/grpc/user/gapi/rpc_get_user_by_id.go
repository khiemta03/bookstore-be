package gapi

import (
	"context"
	"database/sql"

	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/pb"
	"github.com/khiemta03/bookstore-be/user-service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	id := req.GetId()
	convertedId, err := utils.ConvertToUUID(id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, UserNotFoundError)
	}

	user, err := server.store.GetUserByID(ctx, *convertedId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, UserNotFoundError)
		}

		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	res := convertUser(user)

	return res, nil
}
