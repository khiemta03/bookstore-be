package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/authentication/pb"
	errors "github.com/khiemta03/bookstore-be/authentication-service/pkg/error"
)

func (server *Server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	token := req.GetToken()

	payload, err := server.tokenMaker.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf(errors.InvalidAgrumentError)
	}

	dbAccessToken, err := server.store.GetAccessToken(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(errors.InvalidAgrumentError)
		}
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	session, err := server.store.GetSession(ctx, dbAccessToken.SessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(errors.InvalidAgrumentError)
		}
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	if session.Status != "Active" {
		return nil, fmt.Errorf(errors.InvalidAgrumentError)
	}

	if time.Now().After(dbAccessToken.ExpiresAt) {
		return nil, fmt.Errorf(errors.InvalidAgrumentError)
	}

	return &pb.ValidateTokenResponse{UserId: session.UserID}, nil
}
