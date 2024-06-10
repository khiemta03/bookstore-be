package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
)

type CreateOrderTxParams struct {
	UserID          string        `json:"user_id"`
	DiscountID      uuid.NullUUID `json:"discount_id"`
	ShippingAddress string        `json:"shipping_address"`
	CartItemIDList  []uuid.UUID   `json:"cart_items_ids"`
}

type CreateOrderTxResult struct {
	Order           ORDER         `json:"order"`
	Discount        DISCOUNT      `json:"discount"`
	OrderDetailList []ORDERDETAIL `json:"order_detail_list"`
}

// CreateOrderTx handles new order creation transaction
func (store *Store) CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (CreateOrderTxResult, ce.CustomError) {
	var result CreateOrderTxResult

	cerr := store.execTx(ctx, func(q *Queries) ce.CustomError {
		if arg.DiscountID.Valid {
			discount, err := store.GetDiscount(ctx, arg.DiscountID.UUID)
			if err != nil {
				if err == sql.ErrNoRows {
					return ce.DiscountNotFoundError(err)
				}
				return ce.InternalServerError(err)
			}
			result.Discount = discount
		}

		var totalAmount float64 = 0
		orderID := uuid.New()
		for _, cartItemID := range arg.CartItemIDList {
			cartItem, err := store.GetShoppingCartItem(ctx, cartItemID)
			if err != nil {
				if err == sql.ErrNoRows {
					return ce.ShoppingCartItemNotFoundError(err)
				}

				return ce.InternalServerError(err)
			}

			orderDetail, err := store.CreateOrderDetail(ctx, CreateOrderDetailParams{
				OrderID:   orderID,
				BookID:    cartItem.BookID,
				Quantity:  cartItem.Quantity,
				UnitPrice: cartItem.UnitPrice,
			})
			if err != nil {
				return ce.InternalServerError(err)
			}

			totalAmount += cartItem.UnitPrice

			result.OrderDetailList = append(result.OrderDetailList, orderDetail)
		}

		if arg.DiscountID.Valid {
			totalAmount *= result.Discount.DiscountValue
		}
		order, err := store.CreateOrder(ctx, CreateOrderParams{
			OrderID:         orderID,
			UserID:          arg.UserID,
			TotalAmount:     totalAmount,
			ShippingAddress: arg.ShippingAddress,
			Discount:        arg.DiscountID,
		})
		if err != nil {
			return ce.InternalServerError(err)
		}

		result.Order = order

		return ce.NilCustomError()
	})

	return result, cerr
}
