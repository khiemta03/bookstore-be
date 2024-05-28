// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    "USERS" (
        username,
        hashed_password,
        email,
        address,
        full_name,
        phone
    )
VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING id, username, hashed_password, email, phone, full_name, address, status, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	FullName       string `json:"full_name"`
	Phone          string `json:"phone"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (USER, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.Email,
		arg.Address,
		arg.FullName,
		arg.Phone,
	)
	var i USER
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Phone,
		&i.FullName,
		&i.Address,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
    id, username, hashed_password, email, phone, full_name, address, status, created_at
FROM
    "USERS"
WHERE
    "id" = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (USER, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i USER
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Phone,
		&i.FullName,
		&i.Address,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT
    id, username, hashed_password, email, phone, full_name, address, status, created_at
FROM
    "USERS"
WHERE
    "username" = $1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (USER, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, username)
	var i USER
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.Phone,
		&i.FullName,
		&i.Address,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT
    id, username, hashed_password, email, phone, full_name, address, status, created_at
FROM
    "USERS"
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]USER, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []USER
	for rows.Next() {
		var i USER
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.Email,
			&i.Phone,
			&i.FullName,
			&i.Address,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
