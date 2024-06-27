package bookclient

import (
	"context"
	"fmt"
	"io"

	httpError "github.com/khiemta03/bookstore-be/api-getway/internal/error"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/book"
)

func (c *BookServiceClient) GetBook(req *pb.GetBookRequest) (*pb.Book, *httpError.HTTPError) {
	res, err := c.client.GetBook(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *BookServiceClient) ListBooks(req *pb.ListBooksRequest) ([]*pb.Book, *httpError.HTTPError) {
	var responses []*pb.Book
	stream, err := c.client.ListBooks(context.Background(), req)
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
	fmt.Println(responses)
	return responses, nil
}

func (c *BookServiceClient) GetAuthor(req *pb.GetAuthorRequest) (*pb.Author, *httpError.HTTPError) {
	res, err := c.client.GetAuthor(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *BookServiceClient) GetPublisher(req *pb.GetPublisherRequest) (*pb.Publisher, *httpError.HTTPError) {
	res, err := c.client.GetPublisher(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *BookServiceClient) AddNewBook(req *pb.AddNewBookRequest) (*pb.Book, *httpError.HTTPError) {
	res, err := c.client.AddNewBook(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *BookServiceClient) AddNewAuthor(req *pb.AddNewAuthorRequest) (*pb.Author, *httpError.HTTPError) {
	res, err := c.client.AddNewAuthor(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *BookServiceClient) AddNewPublisher(req *pb.AddNewPublisherRequest) (*pb.Publisher, *httpError.HTTPError) {
	res, err := c.client.AddNewPublisher(context.Background(), req)
	if err != nil {
		return nil, httpError.MapGRPCErrorToHTTPError(err)
	}

	return res, nil
}

func (c *BookServiceClient) UpdateBook(req *pb.UpdateBookRequest) *httpError.HTTPError {
	_, err := c.client.UpdateBook(context.Background(), req)
	if err != nil {
		return httpError.MapGRPCErrorToHTTPError(err)
	}

	return nil
}

func (c *BookServiceClient) UpdateAuthor(req *pb.UpdateAuthorRequest) *httpError.HTTPError {
	_, err := c.client.UpdateAuthor(context.Background(), req)
	if err != nil {
		return httpError.MapGRPCErrorToHTTPError(err)
	}

	return nil
}

func (c *BookServiceClient) UpdatePublisher(req *pb.UpdatePublisherRequest) *httpError.HTTPError {
	_, err := c.client.UpdatePublisher(context.Background(), req)
	if err != nil {
		return httpError.MapGRPCErrorToHTTPError(err)
	}

	return nil
}
