package gapi

import (
	"context"
	"fmt"
	"time"

	db "github.com/khiemta03/bookstore-be/authentication-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/authentication/pb"
	user "github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/user/pb"
	errors "github.com/khiemta03/bookstore-be/authentication-service/pkg/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	_, err := server.userServiceClient.ValidateLoginCredentials(&user.ValidateLoginCredentialsRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf(errors.NonGRPCError)
		}

		switch st.Code() {
		case codes.InvalidArgument:
			return nil, fmt.Errorf(errors.InvalidAgrumentError)
		case codes.Internal:
			return nil, fmt.Errorf(errors.InternalServerError)
		default:
			return nil, fmt.Errorf(errors.UnknownGRPCError)
		}
	}

	user, err := server.userServiceClient.GetUserByName(&user.GetUserByNameRequest{
		Username: username,
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

	userId := user.GetId()

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(userId, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(userId, server.config.AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	arg := &db.SessionAndAccessTokenCreationTxParams{
		SessionId:             refreshPayload.ID,
		UserId:                userId,
		UserAgent:             req.GetUserAgent(),
		ClientIp:              req.GetClientIp(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: time.Now().Add(server.config.RefreshTokenDuration),
		AccessTokenId:         accessPayload.ID,
		AccessTokenValue:      accessToken,
		AccessTokenExpiresAt:  time.Now().Add(server.config.AccessTokenDuration),
	}

	result, err := server.store.LoginTx(ctx, *arg)
	if err != nil {
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	res := &pb.LoginResponse{
		Session:     convertSession(result.Session),
		AccessToken: convertToken(result.AccessToken),
	}

	return res, nil
}
