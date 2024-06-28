-- name: CreateSession :one
INSERT INTO
    "SESSIONS" (
        session_id,
        user_id,
        refresh_token,
        user_agent,
        client_ip,
        expires_at
    )
VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetSession :one
SELECT
    *
FROM
    "SESSIONS"
WHERE "session_id" = $1;

