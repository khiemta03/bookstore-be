package api

import (
	"context"

	"github.com/google/uuid"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetOrderByUser(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	userId := req.GetUserId()
	orderId := req.GetOrderId()

	convertedOrderId, err := uuid.Parse(orderId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, ce.ErrOrderNotFoundStr)
	}

	order, err := server.store.GetOrder(ctx, convertedOrderId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, ce.ErrOrderNotFoundStr)
	}

	if order.UserID != userId {
		return nil, status.Errorf(codes.NotFound, ce.ErrOrderNotFoundStr)
	}

	res := convertOrder(order)

	return res, nil
}
