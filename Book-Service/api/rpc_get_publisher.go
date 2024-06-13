package api

import (
	"context"
	"database/sql"

	ce "github.com/khiemta03/bookstore-be/book-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/book-service/internal/grpc/gen/book"
	utils "github.com/khiemta03/bookstore-be/book-service/internal/util"
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
