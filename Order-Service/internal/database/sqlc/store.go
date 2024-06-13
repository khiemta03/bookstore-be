package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/golang/mock/mockgen/model"
	ce "github.com/khiemta03/bookstore-be/order-service/internal/error"
)

type Store interface {
	Querier
	CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (CreateOrderTxResult, ce.CustomError)
	RemoveItemTx(ctx context.Context, arg RemoveItemTxResult) (SHOPPINGCARTITEM, ce.CustomError)
}

// Store provides all functions to execute db queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(dbDriver, dbSource string) (Store, error) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	return &SQLStore{
		db:      conn,
		Queries: New(conn),
	}, nil
}

// execTx excutes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) ce.CustomError) ce.CustomError {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return ce.CustomError{
			OriginalErr: fmt.Errorf("transaction err: %v", err),
			CustomErr:   ce.ErrInternalServer,
		}
	}

	q := New(tx)
	cerr := fn(q)

	if cerr.OriginalErr != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return ce.CustomError{
				OriginalErr: fmt.Errorf("transaction err: %v, rollback err: %v", cerr.OriginalErr, rbErr),
				CustomErr:   cerr.CustomErr,
			}
		}

		return ce.CustomError{
			OriginalErr: fmt.Errorf("req err: %v", cerr.OriginalErr),
			CustomErr:   cerr.CustomErr,
		}
	}

	if cmErr := tx.Commit(); cmErr != nil {
		return ce.CustomError{
			OriginalErr: fmt.Errorf("transaction err: %v, commit err: %v", cerr.OriginalErr, cmErr),
			CustomErr:   ce.ErrInternalServer,
		}
	}

	return ce.NilCustomError()
}
