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

func (server *Server) GetAuthor(ctx context.Context, req *pb.GetAuthorRequest) (*pb.Author, error) {
	id := req.GetId()
	convertedId, err := utils.ConvertToUUID(id)
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.NotFound, ce.ErrAuthorNotFoundStr)
	}

	author, err := server.store.GetAuthor(ctx, *convertedId)
	if err != nil {
		// TODO: send err to log service
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, ce.ErrAuthorNotFoundStr)
		}

		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	res := convertAuthor(author)

	return res, nil
}
