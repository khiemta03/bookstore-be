package api

import (
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	"github.com/khiemta03/bookstore-be/order-service/internal/grpc/client/book"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
)

func (server *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var err error

	convertedDiscountId := uuid.NullUUID{
		Valid: false,
	}
	if req.GetDiscount() != "" {
		convertedDiscountId.UUID, err = uuid.Parse(req.GetDiscount())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, ce.ErrDiscountNotFoundStr)
		}
		convertedDiscountId.Valid = true
	}

	cartItemList, err := server.store.ListShoppingCartItemsByUser(ctx, db.ListShoppingCartItemsByUserParams{
		UserID: req.GetUserId(),
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, ce.ErrShoppingCartItemNotFoundStr)
		}
		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	var cartItemIdList []uuid.UUID
	var convertedCartItemSentToBookService []*book.DecreaseStockQuantityRequest
	for _, item := range cartItemList {
		cartItemIdList = append(cartItemIdList, item.CartItemID)
		convertedCartItemSentToBookService = append(convertedCartItemSentToBookService, &book.DecreaseStockQuantityRequest{
			BookId:   item.BookID,
			Quantity: item.Quantity,
		})
	}

	cartItemList, err = server.store.UpdateShoppingCartItemListStatus(ctx, db.UpdateShoppingCartItemListStatusParams{
		Status:         "ORDERING",
		CartItemIDList: cartItemIdList,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	// TODO: send message to Book Service
	result, err := server.bookServiceClient.DecreaseStockQuantity(ctx, convertedCartItemSentToBookService)
	if err != nil {
		go func() {
			for {
				_, err = server.store.UpdateShoppingCartItemListStatus(ctx, db.UpdateShoppingCartItemListStatusParams{
					Status:         "ADDED",
					CartItemIDList: cartItemIdList,
				})
				if err == nil {
					break
				}
			}
		}()

		if st, ok := status.FromError(err); ok {
			return nil, status.Errorf(st.Code(), st.Message())
		}

		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	fmt.Println(result)

	for i := range cartItemList {
		cartItemList[i].UnitPrice = result[i].UnitPrice
	}

	txResult, cerr := server.store.CreateOrderTx(ctx, db.CreateOrderTxParams{
		UserID:          req.GetUserId(),
		ShippingAddress: req.GetShippingAddress(),
		DiscountID:      convertedDiscountId,
		ItemList:        cartItemList,
	})
	if !cerr.IsNil() {
		go func() {
			for {
				_, err = server.store.UpdateShoppingCartItemListStatus(ctx, db.UpdateShoppingCartItemListStatusParams{
					Status:         "ADDED",
					CartItemIDList: cartItemIdList,
				})
				if err == nil {
					break
				}
			}
		}()
		// TODO: send err to log service
		switch cerr.CustomErr {
		case ce.ErrInternalServer:
			fmt.Println(cerr)
			return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
		default:
			return nil, status.Errorf(codes.InvalidArgument, cerr.CustomErr.Error())
		}
	}
	_, err = server.store.UpdateShoppingCartItemListStatus(ctx, db.UpdateShoppingCartItemListStatusParams{
		Status:         "ORDERED",
		CartItemIDList: cartItemIdList,
	})
	if err != nil {
		go func() {
			for {
				_, err = server.store.UpdateShoppingCartItemListStatus(ctx, db.UpdateShoppingCartItemListStatusParams{
					Status:         "ADDED",
					CartItemIDList: cartItemIdList,
				})
				if err == nil {
					break
				}
			}
		}()
		return nil, status.Errorf(codes.Internal, ce.ErrInternalServerStr)
	}

	res := &pb.CreateOrderResponse{
		Order:           convertOrder(txResult.Order),
		Discount:        convertDiscount(txResult.Discount),
		OrderDetailList: convertOrderDetailList(txResult.OrderDetailList),
	}

	return res, nil
}
