package test

import (
	"time"

	"github.com/google/uuid"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
)

func randomCartItems(bookId string) []db.SHOPPINGCARTITEM {
	cartItems := make([]db.SHOPPINGCARTITEM, 0, 3)
	cartId := uuid.New()
	statuses := []string{"ADDED", "ORDERING", "ORDERED"}
	for _, status := range statuses {
		item := db.SHOPPINGCARTITEM{
			CartItemID: cartId,
			UserID:     "123",
			BookID:     bookId,
			Quantity:   10,
			UnitPrice:  10,
			Status:     status,
			AddedAt:    time.Now(),
		}
		cartItems = append(cartItems, item)
	}
	return cartItems
}

func randomOrder() db.GetOrderRow {
	return db.GetOrderRow{
		OrderID:         uuid.New(),
		UserID:          uuid.NewString(),
		OrderAt:         time.Now(),
		Status:          "PENDING",
		Discount:        uuid.NullUUID{},
		ShippingAddress: "HCM",
		TotalAmount:     1000,
	}
}

func randomOrderDetail(orderId uuid.UUID, item db.SHOPPINGCARTITEM) db.ORDERDETAIL {
	return db.ORDERDETAIL{
		OrderID:   orderId,
		BookID:    item.BookID,
		CreatedAt: time.Now(),
		Quantity:  item.Quantity,
		UnitPrice: item.UnitPrice,
	}
}
