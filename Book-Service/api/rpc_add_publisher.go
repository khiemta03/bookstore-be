package api

import (
	"context"

	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/book-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/book-service/internal/grpc/gen/book"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AddNewPublisher(ctx context.Context, req *pb.AddNewPublisherRequest) (*pb.Publisher, error) {
	publisher, err := server.store.AddNewPublisher(ctx, db.AddNewPublisherParams{
		Name:    req.GetName(),
		Address: convertStringToNullString(req.Address),
	})
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	res := convertPublisher(publisher)

	return res, nil
}
