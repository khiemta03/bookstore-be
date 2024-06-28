package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx excutes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

//

type SessionAndAccessTokenCreationTxParams struct {
	SessionId             uuid.UUID `json:"session_id"`
	UserId                string    `json:"user_id"`
	UserAgent             string    `json:"user_agent"`
	ClientIp              string    `json:"client_ip"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	AccessTokenId         uuid.UUID `json:"access_token_id"`
	AccessTokenValue      string    `json:"access_token_value"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
}

type SessionAndAccessTokenCreationTxResult struct {
	Session     SESSION     `json:"session"`
	AccessToken ACCESSTOKEN `json:"access_token"`
}

func (store *Store) LoginTx(ctx context.Context, arg SessionAndAccessTokenCreationTxParams) (SessionAndAccessTokenCreationTxResult, error) {
	var result SessionAndAccessTokenCreationTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Session, err = store.CreateSession(ctx, CreateSessionParams{
			SessionID:    arg.SessionId,
			UserID:       arg.UserId,
			RefreshToken: arg.RefreshToken,
			UserAgent:    arg.UserAgent,
			ClientIp:     arg.ClientIp,
			ExpiresAt:    arg.RefreshTokenExpiresAt,
		})

		if err != nil {
			return err
		}

		result.AccessToken, err = store.CreateAccessToken(ctx, CreateAccessTokenParams{
			AccessTokenID:    arg.AccessTokenId,
			SessionID:        arg.SessionId,
			AccessTokenValue: arg.AccessTokenValue,
			ExpiresAt:        arg.AccessTokenExpiresAt,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
