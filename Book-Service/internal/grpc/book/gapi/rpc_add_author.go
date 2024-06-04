package gapi

import (
	"context"

	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/book-service/internal/grpc/book/pb"
	ce "github.com/khiemta03/bookstore-be/book-service/pkg/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AddNewAuthor(ctx context.Context, req *pb.AddNewAuthorRequest) (*pb.Author, error) {
	author, err := server.store.AddNewAuthor(ctx, db.AddNewAuthorParams{
		FullName:  req.GetFullName(),
		Birthdate: convertDateToNullTime(req.BirthDate),
	})
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	res := convertAuthor(author)

	return res, nil
}
