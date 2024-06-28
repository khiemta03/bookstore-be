package gapi

import (
	"context"
	"fmt"

	db "github.com/khiemta03/bookstore-be/user-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/pb"
	"github.com/khiemta03/bookstore-be/user-service/pkg/utils"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	AlreadyExistsError = "already exists error"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	password := req.GetPassword()
	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		FullName:       req.Fullname,
		Phone:          req.Phone,
		Address:        req.Address,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, AlreadyExistsError)
			}
		}
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, InternalServerError)
	}

	res := convertUser(user)

	return res, nil
}
