package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/khiemta03/bookstore-be/authentication-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/authentication/pb"
	errors "github.com/khiemta03/bookstore-be/authentication-service/pkg/error"
)

func (server *Server) RenewAccessToken(ctx context.Context, req *pb.RenewAccessTokenRequest) (*pb.RenewAccessTokenResponse, error) {
	refreshToken := req.GetRefreshToken()

	payload, err := server.tokenMaker.ValidateToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf(errors.InvalidAgrumentError)
	}

	session, err := server.store.GetSession(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(errors.InvalidAgrumentError)
		}
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, fmt.Errorf(errors.InvalidAgrumentError)
	}

	if session.Status != "Active" {
		return nil, fmt.Errorf(errors.InvalidAgrumentError)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(session.UserID, server.config.AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	dbAccessToken, err := server.store.CreateAccessToken(ctx, db.CreateAccessTokenParams{
		AccessTokenID:    accessPayload.ID,
		SessionID:        session.SessionID,
		AccessTokenValue: accessToken,
		ExpiresAt:        time.Now().Add(server.config.AccessTokenDuration),
	})
	if err != nil {
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	res := &pb.RenewAccessTokenResponse{
		AccessToken: convertToken(dbAccessToken),
	}

	return res, nil
}
