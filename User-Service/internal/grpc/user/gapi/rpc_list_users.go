package gapi

import (
	"context"
	"database/sql"
	"log"

	db "github.com/khiemta03/bookstore-be/user-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	UserNotFoundError = "user not found error"
)

func (server *Server) ListUsers(req *pb.ListUsersRequest, srv pb.UserService_ListUsersServer) error {
	page := req.GetPage()
	perPage := req.GetPerPage()

	arg := &db.ListUsersParams{
		Limit:  perPage,
		Offset: (page - 1) * perPage,
	}

	userList, err := server.store.ListUsers(context.Background(), *arg)

	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, UserNotFoundError)
		}

		return status.Errorf(codes.Internal, InternalServerError)
	}

	for _, user := range userList {
		convertedUser := convertUser(user)

		if err := srv.Send(convertedUser); err != nil {
			log.Printf("send error %v", err)
		}
	}

	return nil
}
