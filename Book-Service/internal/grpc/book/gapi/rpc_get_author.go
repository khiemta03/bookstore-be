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
