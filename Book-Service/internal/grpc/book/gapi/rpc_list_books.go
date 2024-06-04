package gapi

import (
	"context"
	"log"

	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/book-service/internal/grpc/book/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListBooks(req *pb.ListBooksRequest, srv pb.BookService_ListBooksServer) error {
	page := req.GetPage()
	perPage := req.GetPerPage()

	arg := &db.ListBooksTxParams{
		Limit:  perPage,
		Offset: (page - 1) * perPage,
	}

	listBookTxResult, cerr := server.store.ListBooksTx(context.Background(), *arg)
	if !cerr.IsNil() {
		// TODO: send err to log service
		return status.Errorf(codes.NotFound, cerr.CustomErr.Error())
	}

	for _, bookTxResult := range listBookTxResult {
		res := convertBook(bookTxResult.Book)
		res.Publisher = convertPublisher(bookTxResult.Publisher)
		for _, author := range bookTxResult.Authors {
			res.Authors = append(res.Authors, convertAuthor(author))
		}

		if err := srv.Send(res); err != nil {
			// TODO: send err to log service
			log.Printf("send error %v", err)
		}
	}

	return nil
}
