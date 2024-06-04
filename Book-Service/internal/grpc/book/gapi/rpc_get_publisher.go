package gapi

import (
	"context"
	"database/sql"

	"github.com/khiemta03/bookstore-be/book-service/internal/grpc/book/pb"
	ce "github.com/khiemta03/bookstore-be/book-service/pkg/error"
	utils "github.com/khiemta03/bookstore-be/book-service/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPublisher(ctx context.Context, req *pb.GetPublisherRequest) (*pb.Publisher, error) {
	id := req.GetId()
	convertedId, err := utils.ConvertToUUID(id)
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.NotFound, ce.ErrPublisherNotFoundStr)
	}

	publisher, err := server.store.GetPublisher(ctx, *convertedId)
	if err != nil {
		// TODO: send err to log service
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, ce.ErrPublisherNotFoundStr)
		}

		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	res := convertPublisher(publisher)

	return res, nil
}
