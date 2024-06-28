package api

import (
	"context"

	ce "github.com/khiemta03/bookstore-be/book-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/book-service/internal/grpc/gen/book"
	utils "github.com/khiemta03/bookstore-be/book-service/internal/util"
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
		switch cerr.CustomErr {
		case ce.ErrInternalServer:
			return nil, status.Errorf(codes.NotFound, ce.ErrInternalServerStr)
		default:
			return nil, status.Errorf(codes.InvalidArgument, cerr.CustomErr.Error())
		}
	}

	// format before returning
	res := convertBook(book.Book)
	res.Publisher = convertPublisher(book.Publisher)
	for _, author := range book.Authors {
		res.Authors = append(res.Authors, convertAuthor(author))
	}

	return res, nil
}
