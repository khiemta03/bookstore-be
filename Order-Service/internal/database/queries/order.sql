-- name: GetOrder :one
SELECT
    O.*,  CAST((SUM(OD.quantity * OD.unit_price) * COALESCE(D.discount_value, 1)) AS FLOAT) AS total_amount
FROM
    "ORDERS" AS O LEFT JOIN "ORDER_DETAILS" AS OD 
    ON O.order_id = OD.order_id
    LEFT JOIN "DISCOUNTS" AS D
    ON O.discount = D.discount_id
WHERE O.order_id = $1
GROUP BY
    O.order_id, D.discount_value;

-- name: ListOrders :many
SELECT
    O.*,  CAST((SUM(OD.quantity * OD.unit_price) * COALESCE(D.discount_value, 1)) AS FLOAT) AS total_amount
FROM
    "ORDERS" AS O LEFT JOIN "ORDER_DETAILS" AS OD 
    ON O.order_id = OD.order_id
    LEFT JOIN "DISCOUNTS" AS D
    ON O.discount = D.discount_id
WHERE user_id = $1
GROUP BY
    O.order_id, D.discount_value
LIMIT $2 OFFSET $3;

-- name: CreateOrder :one
INSERT INTO
    "ORDERS" (
        user_id,
        shipping_address,
        discount
    )
VALUES
    ($1, $2, $3) RETURNING *;

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