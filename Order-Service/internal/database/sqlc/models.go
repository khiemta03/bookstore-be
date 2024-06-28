// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type DISCOUNT struct {
	DiscountID    uuid.UUID `json:"discount_id"`
	DiscountCode  string    `json:"discount_code"`
	DiscountValue float64   `json:"discount_value"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	CreatedAt     time.Time `json:"created_at"`
}

type ORDER struct {
	OrderID         uuid.UUID     `json:"order_id"`
	UserID          string        `json:"user_id"`
	OrderAt         time.Time     `json:"order_at"`
	Status          string        `json:"status"`
	Discount        uuid.NullUUID `json:"discount"`
	ShippingAddress string        `json:"shipping_address"`
}

type ORDERDETAIL struct {
	OrderID   uuid.UUID `json:"order_id"`
	BookID    string    `json:"book_id"`
	Quantity  int32     `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	CreatedAt time.Time `json:"created_at"`
}

type SHOPPINGCARTITEM struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
	UserID     string    `json:"user_id"`
	BookID     string    `json:"book_id"`
	Quantity   int32     `json:"quantity"`
	UnitPrice  float64   `json:"unit_price"`
	Status     string    `json:"status"`
	AddedAt    time.Time `json:"added_at"`
}
