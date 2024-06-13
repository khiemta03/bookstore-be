// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddBookAuthor(ctx context.Context, arg AddBookAuthorParams) (BOOKAUTHOR, error)
	AddNewAuthor(ctx context.Context, arg AddNewAuthorParams) (AUTHOR, error)
	AddNewBook(ctx context.Context, arg AddNewBookParams) (BOOK, error)
	AddNewPublisher(ctx context.Context, arg AddNewPublisherParams) (PUBLISHER, error)
	DecreaseStockQuantity(ctx context.Context, arg DecreaseStockQuantityParams) error
	GetAuthor(ctx context.Context, id uuid.UUID) (AUTHOR, error)
	GetBook(ctx context.Context, id uuid.UUID) (BOOK, error)
	GetBookAuthors(ctx context.Context, bookID uuid.UUID) ([]BOOKAUTHOR, error)
	GetPublisher(ctx context.Context, id uuid.UUID) (PUBLISHER, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]BOOK, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) error
}

var _ Querier = (*Queries)(nil)
