-- name: GetPublisher :one
SELECT
    *
FROM
    "PUBLISHERS"
WHERE id = $1;

-- name: AddNewPublisher :one
INSERT INTO
    "PUBLISHERS" (
        name,
        address
    )
VALUES
    ($1, $2) RETURNING *;
