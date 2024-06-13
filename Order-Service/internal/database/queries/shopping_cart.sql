-- name: GetShoppingCartItemByUser :one
SELECT
    *
FROM
    "SHOPPING_CART_ITEMS"
WHERE user_id = $1 and cart_item_id = $2 and status = 'ADDED';

-- name: GetShoppingCartItemByUserForUpdate :one
SELECT
    *
FROM
    "SHOPPING_CART_ITEMS"
WHERE user_id = $1 and cart_item_id = $2 and status = 'ADDED' FOR UPDATE;

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

-- name: UpdateShoppingCartItemListStatus :many
UPDATE
    "SHOPPING_CART_ITEMS" 
SET
    status = $1
WHERE
    cart_item_id = ANY(sqlc.arg(cart_item_id_list)::UUID[]) RETURNING *;

-- name: UpdateShoppingCartItemStatus :one
UPDATE
    "SHOPPING_CART_ITEMS" 
SET
    status = $1
WHERE
    cart_item_id = $2 RETURNING *;


-- name: UpdateShoppingCartItemQuantity :one
UPDATE
    "SHOPPING_CART_ITEMS" 
SET
    quantity = $2
WHERE
    cart_item_id = $1 RETURNING *;

-- name: ListShoppingCartItemsByUser :many
SELECT
    *
FROM
    "SHOPPING_CART_ITEMS"
WHERE user_id = $1
and status = 'ADDED'
LIMIT $2 OFFSET $3;