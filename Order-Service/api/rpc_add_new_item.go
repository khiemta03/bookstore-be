package api

import (
	"context"

	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	"github.com/khiemta03/bookstore-be/order-service/internal/grpc/client/book"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"github.com/lib/pq"
	"google.golang.org/grpc/status"
)

func (server *Server) AddNewItem(ctx context.Context, req *pb.AddNewItemRequest) (*pb.ShoppingCartItem, error) {
	res, err := server.bookServiceClient.CheckBookAdaptability(ctx, &book.CheckBookAdaptabilityRequest{
		BookId:   req.GetBookId(),
		Quantity: req.GetQuantity(),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, ce.ErrNonGRPC
		}

		return nil, status.Errorf(st.Code(), st.Message())
	}

	item, err := server.store.CreateShoppingCartItem(ctx, db.CreateShoppingCartItemParams{
		UserID:    req.GetUserId(),
		BookID:    req.GetBookId(),
		Quantity:  req.GetQuantity(),
		UnitPrice: res.UnitPrice,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, ce.ErrItemExisted
			default:
				return nil, ce.ErrInternalServer
			}
		}
	}

	return convertShoppingCartItem(item), nil
}
