package api

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetItem(ctx context.Context, req *pb.GetShoppingCartItemRequest) (*pb.ShoppingCartItem, error) {
	convertedId, err := uuid.Parse(req.GetItemId())
	if err != nil {
		return nil, status.Error(codes.NotFound, ce.ErrShoppingCartItemNotFoundStr)
	}
	item, err := server.store.GetShoppingCartItemByUser(ctx, db.GetShoppingCartItemByUserParams{
		UserID:     req.GetUserId(),
		CartItemID: convertedId,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, ce.ErrShoppingCartItemNotFoundStr)
		}
		return nil, status.Error(codes.Internal, ce.ErrInternalServerStr)
	}

	return convertShoppingCartItem(item), nil
}
