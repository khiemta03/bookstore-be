package api

import (
	"context"
	"io"

	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/book-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/book-service/internal/grpc/gen/book"
	utils "github.com/khiemta03/bookstore-be/book-service/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DecreaseStockQuantity(stream pb.BookService_DecreaseStockQuantityServer) error {
	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			return status.Errorf(codes.Internal, ce.ErrInternalServerStr)
		}

		quantity := req.GetQuanity()
		if quantity <= 0 {
			return status.Errorf(codes.InvalidArgument, ce.ErrInvalidAgrumentStr)
		}

		id := req.GetBookId()
		convertedId, err := utils.ConvertToUUID(id)
		if err != nil {
			// TODO: send err to log service
			return status.Errorf(codes.NotFound, ce.ErrBookNotFoundStr)
		}

		txResult, cerr := server.store.DecreaseStockQuantityTx(context.Background(), db.DecreaseStockQuantityTxParams{
			BookID:  *convertedId,
			Quanity: req.GetQuanity(),
		})

		if !cerr.IsNil() {
			switch cerr.CustomErr {
			case ce.ErrInternalServer:
				return status.Errorf(codes.Internal, ce.ErrInternalServerStr)
			case ce.ErrOutOfStock:
				return status.Errorf(codes.ResourceExhausted, ce.ErrOutOfStockStr)
			default:
				return status.Errorf(codes.InvalidArgument, cerr.CustomErr.Error())
			}
		}

		res := pb.DecreaseStockQuantityResponse{
			BookId:    id,
			Quantity:  quantity,
			UnitPrice: txResult.UnitPrice,
		}

		if err := stream.Send(&res); err != nil {
			return status.Errorf(codes.Internal, ce.ErrInternalServerStr)
		}
	}

	return nil
}
