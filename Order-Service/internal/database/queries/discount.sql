-- name: GetDiscount :one
SELECT
    *
FROM
    "DISCOUNTS"
WHERE
    discount_id = $1
    AND start_date <= NOW()
    AND end_date >= NOW();

-- name: ListDiscounts :many
SELECT
    *
FROM
    "DISCOUNTS"
WHERE
    start_date <= NOW()
    AND end_date >= NOW()
LIMIT $1 OFFSET $2;

-- name: AddNewDiscount :one
INSERT INTO
    "DISCOUNTS" (
        discount_code,
        discount_value,
        start_date,
        end_date
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;