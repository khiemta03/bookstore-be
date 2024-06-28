package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
)

type CreateOrderTxParams struct {
	UserID          string             `json:"user_id"`
	DiscountID      uuid.NullUUID      `json:"discount_id"`
	ShippingAddress string             `json:"shipping_address"`
	ItemList        []SHOPPINGCARTITEM `json:"item_list"`
}

type CreateOrderTxResult struct {
	Order           GetOrderRow   `json:"order"`
	Discount        DISCOUNT      `json:"discount"`
	OrderDetailList []ORDERDETAIL `json:"order_detail_list"`
}

// CreateOrderTx handles new order creation transaction
func (store *SQLStore) CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (CreateOrderTxResult, ce.CustomError) {
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

		order, err := store.CreateOrder(ctx, CreateOrderParams{
			UserID:          arg.UserID,
			ShippingAddress: arg.ShippingAddress,
			Discount:        arg.DiscountID,
		})
		if err != nil {
			return ce.InternalServerError(err)
		}

		for _, cartItem := range arg.ItemList {
			orderDetail, err := store.CreateOrderDetail(ctx, CreateOrderDetailParams{
				OrderID:   order.OrderID,
				BookID:    cartItem.BookID,
				Quantity:  cartItem.Quantity,
				UnitPrice: cartItem.UnitPrice,
			})

			if err != nil {
				return ce.InternalServerError(err)
			}

			result.OrderDetailList = append(result.OrderDetailList, orderDetail)
		}

		orderWithTotalAmount, err := store.GetOrder(ctx, order.OrderID)
		if err != nil {
			return ce.InternalServerError(err)
		}

		result.Order = orderWithTotalAmount

		return ce.NilCustomError()
	})

	return result, cerr
}

type RemoveItemTxResult struct {
	UserID string    `json:"user_id"`
	ItemID uuid.UUID `json:"item_id"`
}

func (store *SQLStore) RemoveItemTx(ctx context.Context, arg RemoveItemTxResult) (SHOPPINGCARTITEM, ce.CustomError) {
	var result SHOPPINGCARTITEM
	cerr := store.execTx(ctx, func(q *Queries) ce.CustomError {
		_, err := store.GetShoppingCartItemByUserForUpdate(ctx, GetShoppingCartItemByUserForUpdateParams{
			UserID:     arg.UserID,
			CartItemID: arg.ItemID,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return ce.ShoppingCartItemNotFoundError(err)
			}
			return ce.InternalServerError(err)
		}

		result, err = store.UpdateShoppingCartItemStatus(ctx, UpdateShoppingCartItemStatusParams{
			CartItemID: arg.ItemID,
			Status:     "REMOVED",
		})
		if err != nil {
			return ce.InternalServerError(err)
		}

		return ce.NilCustomError()
	})

	return result, cerr
}
