package book

import (
	"context"

	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/book"
)

type CheckBookAdaptabilityRequest struct {
	BookId   string `json:"book_id"`
	Quantity int32  `json:"quantity"`
}

type CheckBookAdaptabilityResponse struct {
	BookId    string  `json:"book_id"`
	Quantity  int32   `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

func (c *BookServiceClient) CheckBookAdaptability(ctx context.Context, req *CheckBookAdaptabilityRequest) (*CheckBookAdaptabilityResponse, error) {
	res, err := c.client.CheckBookAdaptability(ctx, &pb.CheckBookAdaptabilityRequest{
		BookId:   req.BookId,
		Quantity: req.Quantity,
	})
	if err != nil {
		return nil, err
	}

	return &CheckBookAdaptabilityResponse{
		BookId:    res.GetBookId(),
		Quantity:  res.GetQuantity(),
		UnitPrice: res.GetUnitPrice(),
	}, nil
}
