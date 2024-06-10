-- name: GetOrder :one
SELECT
    *
FROM
    "ORDERS"
WHERE order_id = $1;

-- name: ListOrdersByUser :one
SELECT
    *
FROM
    "ORDERS"
WHERE user_id = $1
LIMIT $2 OFFSET $3;

-- name: ListOrders :many
SELECT
    *
FROM
    "ORDERS"
LIMIT $1 OFFSET $2;

-- name: CreateOrder :one
INSERT INTO
    "ORDERS" (
        order_id,
        user_id,
        total_amount,
        shipping_address,
        discount
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;

-- name: ListOrderDetails :many
SELECT
    *
FROM
    "ORDER_DETAILS"
WHERE order_id = $1
LIMIT $2 OFFSET $3;

-- name: GetOrderDetail :one
SELECT
    *
FROM
    "ORDER_DETAILS"
WHERE order_id = $1 and book_id = $2;

-- name: CreateOrderDetail :one
INSERT INTO
    "ORDER_DETAILS" (
        order_id,
        book_id,
        quantity,
        unit_price
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;