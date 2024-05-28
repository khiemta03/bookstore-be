package gapi

import (
	db "github.com/khiemta03/bookstore-be/user-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.USER) *pb.User {

	return &pb.User{
		Id:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Fullname:  user.FullName,
		Phone:     user.Phone,
		Address:   user.Address,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
