package orderclient

import (
	"context"
	"io"

	httpError "github.com/khiemta03/bookstore-be/api-getway/internal/error"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/order"
)

func (c *OrderServiceClient) GetOrderByUser(req *pb.GetOrderRequest) (*pb.Order, *httpError.HTTPError) {
	res, err := c.client.GetOrderByUser(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *OrderServiceClient) ListItems(req *pb.ListShoppingCartItemsRequest) ([]*pb.ShoppingCartItem, *httpError.HTTPError) {
	var responses []*pb.ShoppingCartItem
	stream, err := c.client.ListItems(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, httpError.MapGRPCErrorToHTTPError(err)
		}
		responses = append(responses, res)
	}

	return responses, nil
}

func (c *OrderServiceClient) CreateOrder(req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, *httpError.HTTPError) {
	res, err := c.client.CreateOrder(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *OrderServiceClient) AddNewItem(req *pb.AddNewItemRequest) (*pb.ShoppingCartItem, *httpError.HTTPError) {
	res, err := c.client.AddNewItem(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *OrderServiceClient) RemoveItem(req *pb.RemoveItemRequest) (*pb.RemoveItemResponse, *httpError.HTTPError) {
	res, err := c.client.RemoveItem(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *OrderServiceClient) GetItem(req *pb.GetShoppingCartItemRequest) (*pb.ShoppingCartItem, *httpError.HTTPError) {
	res, err := c.client.GetItem(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}
