-- name: ListBooks :many
SELECT
    *
FROM
    "BOOKS"
LIMIT $1 OFFSET $2;

-- name: GetBook :one
SELECT
    *
FROM
    "BOOKS"
WHERE id = $1;

-- name: GetBookAuthors :many
SELECT
    *
FROM
    "BOOK_AUTHORS"
WHERE book_id = $1;

-- name: AddBookAuthor :one
INSERT INTO
    "BOOK_AUTHORS" (
        book_id,
        author_id
    )
VALUES
    ($1, $2) RETURNING *;


-- name: AddNewBook :one
INSERT INTO
    "BOOKS" (
        title,
        full_title,
        publisher,
        publication_date,
        isbn,
        description,
        price,
        stock_quantity,
        front_cover_image,
        back_cover_image
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;

-- name: UpdateBook :exec
UPDATE
    "BOOKS"
SET
    title               = COALESCE($2, title),
    full_title          = COALESCE($3, full_title), 
    publisher           = COALESCE($4, publisher),
    publication_date    = COALESCE($5, publication_date),
    isbn                = COALESCE($6, isbn),
    description         = COALESCE($7, description),
    price               = COALESCE($8, price),
    stock_quantity      = COALESCE($9, stock_quantity),
    front_cover_image   = COALESCE($10, front_cover_image),
    back_cover_image    = COALESCE($11, back_cover_image)
WHERE
    id = $1;

-- name: DecreaseStockQuantity :exec
UPDATE
    "BOOKS"
SET
    stock_quantity = stock_quantity - sqlc.arg(quantity)
WHERE
    id = $1;
    

