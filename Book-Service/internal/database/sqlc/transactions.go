package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	ce "github.com/khiemta03/bookstore-be/book-service/internal/error"
	"github.com/lib/pq"
)

type AddNewBookTxParams struct {
	Title           string         `json:"title"`
	FullTitle       string         `json:"full_title"`
	Publisher       uuid.UUID      `json:"publisher"`
	PublicationDate time.Time      `json:"publication_date"`
	Isbn            string         `json:"isbn"`
	Description     sql.NullString `json:"description"`
	Price           float64        `json:"price"`
	StockQuantity   int32          `json:"stock_quantity"`
	FrontCoverImage sql.NullString `json:"front_cover_image"`
	BackCoverImage  sql.NullString `json:"back_cover_image"`
	Authors         []uuid.UUID    `json:"authors"`
}

type AddNewBookTxResult struct {
	Book      BOOK      `json:"book"`
	Publisher PUBLISHER `json:"publisher"`
	Authors   []AUTHOR  `json:"authors"`
}

// AddNewBookTx handles adding new book transaction
func (store *Store) AddNewBookTx(ctx context.Context, arg AddNewBookTxParams) (AddNewBookTxResult, ce.CustomError) {
	var result AddNewBookTxResult

	cerr := store.execTx(ctx, func(q *Queries) ce.CustomError {
		publisher, err := q.GetPublisher(ctx, arg.Publisher)
		if err != nil {
			if err == sql.ErrNoRows {
				return ce.PublisherNotFoundError(err)
			}

			return ce.InternalServerError(err)
		}
		result.Publisher = publisher

		book, err := q.AddNewBook(ctx, AddNewBookParams{
			Title:           arg.Title,
			FullTitle:       arg.FullTitle,
			Publisher:       arg.Publisher,
			PublicationDate: arg.PublicationDate,
			Isbn:            arg.Isbn,
			Description:     arg.Description,
			Price:           arg.Price,
			StockQuantity:   arg.StockQuantity,
			FrontCoverImage: arg.FrontCoverImage,
			BackCoverImage:  arg.BackCoverImage,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				log.Println(pqErr.Code.Name())
			}

			return ce.InternalServerError(err)
		}

		// check if author has already existed
		for _, id := range arg.Authors {
			author, err := q.GetAuthor(ctx, id)
			if err != nil {
				if err == sql.ErrNoRows {
					return ce.AuthorNotFoundError(err)
				}

				return ce.InternalServerError(err)
			}

			_, err = q.AddBookAuthor(ctx, AddBookAuthorParams{
				BookID:   book.ID,
				AuthorID: id,
			})
			if err != nil {
				return ce.InternalServerError(err)
			}

			result.Authors = append(result.Authors, author)
		}

		result.Book = book

		return ce.NilCustomError()
	})

	return result, cerr
}

type GetBookTxResult struct {
	Book      BOOK      `json:"book"`
	Publisher PUBLISHER `json:"publisher"`
	Authors   []AUTHOR  `json:"authors"`
}

// GetBookTx handles getting a specific book transaction
func (store *Store) GetBookTx(ctx context.Context, id uuid.UUID) (GetBookTxResult, ce.CustomError) {
	var result GetBookTxResult

	cerr := store.execTx(ctx, func(q *Queries) ce.CustomError {
		book, err := q.GetBook(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return ce.BookNotFoundError(err)
			}

			return ce.InternalServerError(err)
		}
		result.Book = book

		publisher, err := q.GetPublisher(ctx, book.Publisher)
		if err != nil {
			if err == sql.ErrNoRows {
				return ce.PublisherNotFoundError(err)
			}

			return ce.InternalServerError(err)
		}
		result.Publisher = publisher

		authors, err := q.GetBookAuthors(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return ce.AuthorNotFoundError(err)
			}

			return ce.InternalServerError(err)
		}

		for _, author := range authors {
			author, err := q.GetAuthor(ctx, author.AuthorID)
			if err != nil {
				if err == sql.ErrNoRows {
					return ce.AuthorNotFoundError(err)
				}

				return ce.InternalServerError(err)
			}
			result.Authors = append(result.Authors, author)
		}

		return ce.NilCustomError()
	})

	return result, cerr
}

type ListBooksTxParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

// ListBooksTx handles a transaction of getting book list
func (store *Store) ListBooksTx(ctx context.Context, arg ListBooksTxParams) ([]GetBookTxResult, ce.CustomError) {
	var result []GetBookTxResult

	cerr := store.execTx(ctx, func(q *Queries) ce.CustomError {
		books, err := q.ListBooks(ctx, ListBooksParams{
			Offset: arg.Offset,
			Limit:  arg.Limit,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return ce.BookNotFoundError(err)
			}

			return ce.InternalServerError(err)
		}

		for _, book := range books {
			bookTxResult := GetBookTxResult{
				Book: book,
			}

			publisher, err := q.GetPublisher(ctx, book.Publisher)
			if err != nil {
				if err == sql.ErrNoRows {
					return ce.PublisherNotFoundError(err)
				}

				return ce.InternalServerError(err)
			}
			bookTxResult.Publisher = publisher

			authors, err := q.GetBookAuthors(ctx, book.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					return ce.AuthorNotFoundError(err)
				}

				return ce.InternalServerError(err)
			}

			for _, author := range authors {
				author, err := q.GetAuthor(ctx, author.AuthorID)
				if err != nil {
					if err == sql.ErrNoRows {
						return ce.AuthorNotFoundError(err)
					}

					return ce.InternalServerError(err)
				}
				bookTxResult.Authors = append(bookTxResult.Authors, author)
			}

			result = append(result, bookTxResult)
		}

		return ce.NilCustomError()
	})

	return result, cerr
}

type DecreaseStockQuantityTxParams struct {
	BookID  uuid.UUID `json:"user_id"`
	Quanity int32     `json:"quantity"`
}

type DecreaseStockQuantityTxResult struct {
	BookID    uuid.UUID `json:"user_id"`
	Quanity   int32     `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
}

// ListBooksTx handles a transaction of decreasing a book's stock quantity
func (store *Store) DecreaseStockQuantityTx(ctx context.Context, arg DecreaseStockQuantityTxParams) (*DecreaseStockQuantityTxResult, ce.CustomError) {
	var result DecreaseStockQuantityTxResult

	cerr := store.execTx(ctx, func(q *Queries) ce.CustomError {
		book, err := store.GetBook(ctx, arg.BookID)
		if err != nil {
			if err == sql.ErrNoRows {
				return ce.BookNotFoundError(err)
			}
			return ce.InternalServerError(err)
		}

		if book.StockQuantity < arg.Quanity {
			return ce.OutOfStockError(err)
		}

		err = store.DecreaseStockQuantity(ctx, DecreaseStockQuantityParams{
			ID:       arg.BookID,
			Quantity: arg.Quanity,
		})
		if err != nil {
			return ce.InternalServerError(err)
		}

		result.BookID = arg.BookID
		result.Quanity = arg.Quanity
		result.UnitPrice = book.Price

		return ce.NilCustomError()
	})
	return &result, cerr
}
