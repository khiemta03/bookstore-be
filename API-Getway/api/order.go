package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/order"
)

type getOrderRequest struct {
	OrderId string `uri:"id" binding:"required"`
}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userId, exists := ctx.Get("UserId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("UserId not found in context")))
		return
	}

	userIdString, _ := userId.(string)

	order, err := server.client.GetOrderByUser(&pb.GetOrderRequest{
		OrderId: req.OrderId,
		UserId:  userIdString,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

type getItemRequest struct {
	ItemId string `uri:"id" binding:"required"`
}

func (server *Server) getItem(ctx *gin.Context) {
	var req getItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userId, exists := ctx.Get("UserId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("UserId not found in context")))
		return
	}

	userIdString, _ := userId.(string)

	item, err := server.client.GetItem(&pb.GetShoppingCartItemRequest{
		ItemId: req.ItemId,
		UserId: userIdString,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, item)
}

type addItemRequest struct {
	BookId   string `json:"book_id" binding:"required"`
	Quantity int32  `json:"quantity" binding:"required"`
}

func (server *Server) addItem(ctx *gin.Context) {
	var req addItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userId, exists := ctx.Get("UserId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("UserId not found in context")))
		return
	}

	userIdString, _ := userId.(string)

	item, err := server.client.AddNewItem(&pb.AddNewItemRequest{
		UserId:   userIdString,
		BookId:   req.BookId,
		Quantity: req.Quantity,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, item)
}

type removeItemRequest struct {
	CartItemId string `uri:"id" binding:"required"`
}

func (server *Server) removeItem(ctx *gin.Context) {
	var req removeItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userId, exists := ctx.Get("UserId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("UserId not found in context")))
		return
	}

	userIdString, _ := userId.(string)

	res, err := server.client.RemoveItem(&pb.RemoveItemRequest{
		UserId:     userIdString,
		CartItemId: req.CartItemId,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type listItemsRequest struct {
	Page    int32 `form:"page,min=1,default=1" binding:"required"`
	PerPage int32 `form:"per_page,min=1,default=10" binding:"required"`
}

func (server *Server) listItems(ctx *gin.Context) {
	var req listItemsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userId, exists := ctx.Get("UserId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("UserId not found in context")))
		return
	}

	userIdString, _ := userId.(string)

	res, err := server.client.ListItems(&pb.ListShoppingCartItemsRequest{
		UserId:  userIdString,
		Page:    req.Page,
		PerPage: req.PerPage,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type createOrderRequest struct {
	ShippingAddress string `json:"shipping_address" binding:"required"`
	Discount        string `json:"discount"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userId, exists := ctx.Get("UserId")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("UserId not found in context")))
		return
	}

	userIdString, _ := userId.(string)

	order, err := server.client.CreateOrder(&pb.CreateOrderRequest{
		UserId:          userIdString,
		Discount:        &req.Discount,
		ShippingAddress: req.ShippingAddress,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, order)
}
