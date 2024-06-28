package api

import (
	"context"
	"database/sql"
	"log"

	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListItems(req *pb.ListShoppingCartItemsRequest, srv pb.OrderService_ListItemsServer) error {
	page := req.GetPage()
	perPage := req.GetPerPage()

	arg := db.ListShoppingCartItemsByUserParams{
		UserID: req.UserId,
		Limit:  perPage,
		Offset: (page - 1) * perPage,
	}

	itemList, err := server.store.ListShoppingCartItemsByUser(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Error(codes.NotFound, ce.ErrShoppingCartItemNotFoundStr)
		}
		return status.Error(codes.Internal, ce.ErrInternalServerStr)
	}

	for _, item := range itemList {
		convertedItem := convertShoppingCartItem(item)
		if err := srv.Send(convertedItem); err != nil {
			log.Printf("send error %v", err)
		}
	}

	return nil
}
