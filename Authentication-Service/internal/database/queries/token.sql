-- name: CreateAccessToken :one
INSERT INTO
    "ACCESS_TOKEN" (
        access_token_id,
        session_id,
        access_token_value,
        expires_at
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetAccessToken :one
SELECT
    *
FROM
    "ACCESS_TOKEN"
WHERE "access_token_id" = $1;

