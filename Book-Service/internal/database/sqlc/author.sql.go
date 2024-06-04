// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: author.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addNewAuthor = `-- name: AddNewAuthor :one
INSERT INTO
    "AUTHORS" (
        full_name,
        birthdate
    )
VALUES
    ($1, $2) RETURNING id, full_name, birthdate, created_at
`

type AddNewAuthorParams struct {
	FullName  string       `json:"full_name"`
	Birthdate sql.NullTime `json:"birthdate"`
}

func (q *Queries) AddNewAuthor(ctx context.Context, arg AddNewAuthorParams) (AUTHOR, error) {
	row := q.db.QueryRowContext(ctx, addNewAuthor, arg.FullName, arg.Birthdate)
	var i AUTHOR
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Birthdate,
		&i.CreatedAt,
	)
	return i, err
}

const getAuthor = `-- name: GetAuthor :one
SELECT
    id, full_name, birthdate, created_at
FROM
    "AUTHORS"
WHERE id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, id uuid.UUID) (AUTHOR, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i AUTHOR
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Birthdate,
		&i.CreatedAt,
	)
	return i, err
}
