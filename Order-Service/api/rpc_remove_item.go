package api

import (
	"context"

	"github.com/google/uuid"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RemoveItem(ctx context.Context, req *pb.RemoveItemRequest) (*pb.RemoveItemResponse, error) {
	convertedItemID, err := uuid.Parse(req.GetCartItemId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, ce.ErrShoppingCartItemNotFoundStr)
	}

	_, cerr := server.store.RemoveItemTx(ctx, db.RemoveItemTxResult{
		UserID: req.GetUserId(),
		ItemID: convertedItemID,
	})
	if !cerr.IsNil() {
		switch cerr.CustomErr {
		case ce.ErrInternalServer:
			return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
		default:
			return nil, status.Errorf(codes.InvalidArgument, cerr.CustomErr.Error())
		}
	}

	return &pb.RemoveItemResponse{
		IsSuccessful: true,
	}, nil
}
