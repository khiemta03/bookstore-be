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

func (server *Server) CheckBookAdaptability(ctx context.Context, req *pb.CheckBookAdaptabilityRequest) (*pb.CheckBookAdaptabilityResponse, error) {
	quantity := req.GetQuantity()
	if quantity <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, ce.ErrInvalidAgrumentStr)
	}

	id := req.GetBookId()
	convertedId, err := utils.ConvertToUUID(id)
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.NotFound, ce.ErrBookNotFoundStr)
	}

	book, err := server.store.GetBook(ctx, *convertedId)
	if err != nil {
		// TODO: send err to log service
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, ce.ErrBookNotFoundStr)
		}
		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	if book.StockQuantity < quantity {
		return nil, status.Errorf(codes.ResourceExhausted, ce.ErrOutOfStockStr)
	}

	return &pb.CheckBookAdaptabilityResponse{
		BookId:    id,
		Quantity:  quantity,
		UnitPrice: book.Price,
	}, nil
}
