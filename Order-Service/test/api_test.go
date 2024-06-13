package test

import (
	"context"
	"database/sql"
	"log"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/khiemta03/bookstore-be/order-service/api"
	mock_sqlc "github.com/khiemta03/bookstore-be/order-service/internal/database/mock"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
	"github.com/khiemta03/bookstore-be/order-service/internal/grpc/client/book"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func TestAPI_GetOrder(t *testing.T) {
	order := randomOrder()
	invalidId := uuid.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_sqlc.NewMockStore(ctrl)
	store.EXPECT().
		GetOrder(gomock.Any(), gomock.Eq(order.OrderID)).
		Times(1).
		Return(order, nil)

	store.EXPECT().
		GetOrder(gomock.Any(), gomock.Not(order.OrderID)).
		Times(1).
		Return(db.GetOrderRow{}, sql.ErrNoRows)

	listener := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		listener.Close()
	})

	server := api.NewMockServer(store, nil)

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, server)
	go func() {
		err := s.Serve(listener)
		log.Fatalf("Errof Serve:%v", err)
		require.NoError(t, err)
	}()

	//TEST
	dialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := pb.NewOrderServiceClient(conn)
	// Test case 1: Success
	res1, err := client.GetOrderByUser(context.Background(), &pb.GetOrderRequest{
		OrderId: order.OrderID.String(),
		UserId:  order.UserID,
	})
	require.NoError(t, err)
	require.Equal(t, res1.OrderId, order.OrderID.String())

	// Test case 2: Failure - Invalid Order Id
	res2, err := client.GetOrderByUser(context.Background(), &pb.GetOrderRequest{
		OrderId: invalidId.String(),
		UserId:  order.UserID,
	})
	require.Error(t, err)
	require.Nil(t, res2)
	if st, ok := status.FromError(err); ok {
		require.Equal(t, st.Code(), codes.NotFound)
	}
}

func TestAPI_CreateOrder(t *testing.T) {
	// prepare data
	order := randomOrder()
	var originalCartItemList []db.SHOPPINGCARTITEM
	var orderingCartItemList []db.SHOPPINGCARTITEM
	var orderedCartItemList []db.SHOPPINGCARTITEM
	var cartItemIdList []uuid.UUID
	var orderDetailList []db.ORDERDETAIL

	bookIdList := [...]string{"5a96b419-5902-48f0-934e-57b066c8f69c", "58f2f4fb-e64a-4858-b73b-51447b646bf1", "6f3c5b01-6913-443e-850f-44736c48a8bf"}
	for _, bookId := range bookIdList {
		items := randomCartItems(bookId)
		cartItemIdList = append(cartItemIdList, items[0].CartItemID)
		originalCartItemList = append(originalCartItemList, items[0])
		orderingCartItemList = append(orderingCartItemList, items[1])
		orderedCartItemList = append(orderedCartItemList, items[2])
		orderDetailList = append(orderDetailList, randomOrderDetail(order.OrderID, items[0]))
	}

	txResult := db.CreateOrderTxResult{
		Order:           order,
		Discount:        db.DISCOUNT{},
		OrderDetailList: orderDetailList,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_sqlc.NewMockStore(ctrl)
	store.EXPECT().
		CreateOrderTx(gomock.Any(), gomock.Eq(db.CreateOrderTxParams{
			UserID:          order.UserID,
			ShippingAddress: order.ShippingAddress,
			DiscountID:      order.Discount,
			ItemList:        orderingCartItemList,
		})).
		AnyTimes().
		Return(txResult, ce.NilCustomError())

	store.EXPECT().
		ListShoppingCartItemsByUser(gomock.Any(), gomock.Eq(db.ListShoppingCartItemsByUserParams{
			UserID: order.UserID,
			Limit:  10,
			Offset: 0,
		})).
		AnyTimes().
		Return(originalCartItemList, nil)

	store.EXPECT().
		UpdateShoppingCartItemListStatus(gomock.Any(), gomock.Eq(db.UpdateShoppingCartItemListStatusParams{
			Status:         "ORDERING",
			CartItemIDList: cartItemIdList,
		})).
		Return(orderingCartItemList, nil)

	store.EXPECT().
		UpdateShoppingCartItemListStatus(gomock.Any(), gomock.Eq(db.UpdateShoppingCartItemListStatusParams{
			Status:         "ADDED",
			CartItemIDList: cartItemIdList,
		})).
		AnyTimes().
		Return(originalCartItemList, nil)

	store.EXPECT().
		UpdateShoppingCartItemListStatus(gomock.Any(), gomock.Eq(db.UpdateShoppingCartItemListStatusParams{
			Status:         "ORDERED",
			CartItemIDList: cartItemIdList,
		})).
		AnyTimes().
		Return(orderedCartItemList, nil)

	listener := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		listener.Close()
	})

	server := api.NewMockServer(store, book.NewBookServiceClient("localhost:3003"))

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, server)
	go func() {
		err := s.Serve(listener)
		log.Fatalf("Errof Serve:%v", err)
		require.NoError(t, err)
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Close()
	})

	client := pb.NewOrderServiceClient(conn)
	// Test case 1: Success
	res1, err := client.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		UserId:          order.UserID,
		Discount:        nil,
		ShippingAddress: "HCM",
	})
	require.NoError(t, err)
	require.Equal(t, res1.Order.Status, "PENDING")
}
