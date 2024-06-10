package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
)

func (server *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var convertedCartItemIdList []uuid.UUID
	var err error
	for _, id := range req.CartItemIdList {
		convertedId, err := uuid.Parse(id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, ce.ErrShoppingCartItemNotFoundStr)
		}
		convertedCartItemIdList = append(convertedCartItemIdList, convertedId)
	}

	convertedDiscountId := uuid.NullUUID{
		Valid: false,
	}
	if req.Discount != nil {
		convertedDiscountId.UUID, err = uuid.Parse(*req.Discount)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, ce.ErrDiscountNotFoundStr)
		}
		convertedDiscountId.Valid = true
	}
	txResult, cerr := server.store.CreateOrderTx(ctx, db.CreateOrderTxParams{
		UserID:         req.GetUserId(),
		DiscountID:     convertedDiscountId,
		CartItemIDList: convertedCartItemIdList,
	})
	if !cerr.IsNil() {
		// TODO: send err to log service
		switch cerr.CustomErr {
		case ce.ErrInternalServer:
			return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
		default:
			return nil, status.Errorf(codes.InvalidArgument, cerr.CustomErr.Error())
		}
	}

	res := &pb.CreateOrderResponse{
		Order:           convertOrder(txResult.Order),
		Discount:        convertDiscount(txResult.Discount),
		OrderDetailList: convertOrderDetailList(txResult.OrderDetailList),
	}

	return res, nil
}
