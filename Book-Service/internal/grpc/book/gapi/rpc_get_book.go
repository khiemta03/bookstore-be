package gapi

import (
	"context"

	"github.com/khiemta03/bookstore-be/book-service/internal/grpc/book/pb"
	ce "github.com/khiemta03/bookstore-be/book-service/pkg/error"
	utils "github.com/khiemta03/bookstore-be/book-service/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
	id := req.GetId()
	convertedId, err := utils.ConvertToUUID(id)
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.NotFound, ce.ErrBookNotFoundStr)
	}

	book, cerr := server.store.GetBookTx(ctx, *convertedId)
	if !cerr.IsNil() {
		// TODO: send err to log service
		return nil, status.Errorf(codes.NotFound, cerr.CustomErr.Error())
	}

	// format before returning
	res := convertBook(book.Book)
	res.Publisher = convertPublisher(book.Publisher)
	for _, author := range book.Authors {
		res.Authors = append(res.Authors, convertAuthor(author))
	}

	return res, nil
}
