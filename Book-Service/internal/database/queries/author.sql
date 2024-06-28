-- name: GetAuthor :one
SELECT
    *
FROM
    "AUTHORS"
WHERE id = $1;

-- name: AddNewAuthor :one
INSERT INTO
    "AUTHORS" (
        full_name,
        birthdate
    )
VALUES
    ($1, $2) RETURNING *;
