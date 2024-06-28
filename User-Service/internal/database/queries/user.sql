-- name: CreateUser :one
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
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: ListUsers :many
SELECT
    *
FROM
    "USERS"
LIMIT $1 OFFSET $2;

-- name: GetUserByID :one
SELECT
    *
FROM
    "USERS"
WHERE
    "id" = $1;

-- name: GetUserByName :one
SELECT
    *
FROM
    "USERS"
WHERE
    "username" = $1;