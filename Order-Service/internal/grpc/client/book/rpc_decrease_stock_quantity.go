package book

import (
	"context"
	"fmt"
	"io"

	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/book"
)

type DecreaseStockQuantityRequest struct {
	BookId   string `json:"book_id"`
	Quantity int32  `json:"quantity"`
}

type DecreaseStockQuantityResponse struct {
	BookId    string  `json:"book_id"`
	Quantity  int32   `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

func (c *BookServiceClient) DecreaseStockQuantity(ctx context.Context, requests []*DecreaseStockQuantityRequest) ([]*DecreaseStockQuantityResponse, error) {
	var responses []*DecreaseStockQuantityResponse
	stream, err := c.client.DecreaseStockQuantity(ctx)
	if err != nil {
		return nil, err
	}

	// Send each request to the server
	for _, req := range requests {
		if err := stream.Send(&pb.DecreaseStockQuantityRequest{
			BookId:   req.BookId,
			Quantity: req.Quantity,
		}); err != nil {
			// return nil, err
			fmt.Println(err)
		}
	}

	stream.CloseSend()

	// Receive responses from the server
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		responses = append(responses, &DecreaseStockQuantityResponse{
			BookId:    res.GetBookId(),
			Quantity:  res.GetQuantity(),
			UnitPrice: res.GetUnitPrice(),
		})
	}

	return responses, nil
}
