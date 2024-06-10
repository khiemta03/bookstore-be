-- name: GetShoppingCartItem :one
SELECT
    *
FROM
    "SHOPPING_CART_ITEMS"
WHERE cart_item_id = $1;

-- name: CreateShoppingCartItem :one
INSERT INTO
    "SHOPPING_CART_ITEMS" (
        user_id,
        book_id,
        quantity,
        unit_price
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: ListShoppingCartItemsByUser :many
SELECT
    *
FROM
    "SHOPPING_CART_ITEMS"
WHERE user_id = $1
LIMIT $1 OFFSET $2;