package gapi

import (
	db "github.com/khiemta03/bookstore-be/authentication-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/authentication/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertSession(session db.SESSION) *pb.Session {
	return &pb.Session{
		SessionId:    session.SessionID.String(),
		RefreshToken: session.RefreshToken,
		ExpiresAt:    timestamppb.New(session.ExpiresAt),
	}
}

func convertToken(token db.ACCESSTOKEN) *pb.AccessToken {
	return &pb.AccessToken{
		AccessTokenValue: token.AccessTokenValue,
		ExpiresAt:        timestamppb.New(token.ExpiresAt),
	}
}
