package gapi

import (
	"context"
	"database/sql"

	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/pb"
	"github.com/khiemta03/bookstore-be/user-service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidCredentialsError = "invalid credentials error"
	InternalServerError     = "internal server error"
	ValidCredentials        = "valid credentials"

	InvalidCredentialsResponse = &pb.ValidateLoginCredentialsResponse{
		IsValid: false,
		Error:   InvalidCredentialsError,
	}

	InternalCredentialsResponse = &pb.ValidateLoginCredentialsResponse{
		IsValid: false,
		Error:   InternalServerError,
	}

	ValidCredentialsResponse = &pb.ValidateLoginCredentialsResponse{
		IsValid: true,
		Error:   ValidCredentials,
	}
)

func (server *Server) ValidateLoginCredentials(ctx context.Context, req *pb.ValidateLoginCredentialsRequest) (*pb.ValidateLoginCredentialsResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	user, err := server.store.GetUserByName(ctx, username)

	if err != nil {
		if err == sql.ErrNoRows {
			return InvalidCredentialsResponse, status.Errorf(codes.InvalidArgument, InvalidCredentialsError)
		}

		return InternalCredentialsResponse, status.Errorf(codes.Internal, InternalServerError)
	}

	err = utils.CheckPassword(password, user.HashedPassword)

	if err != nil {
		return InvalidCredentialsResponse, status.Errorf(codes.InvalidArgument, InvalidCredentialsError)
	}

	return ValidCredentialsResponse, nil
}
