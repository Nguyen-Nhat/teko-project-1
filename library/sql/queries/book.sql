-- name: GetAllBooks :many
SELECT id, title, published_year, available_quantity, borrow_quantity, created_at, updated_at
FROM "Book";

-- name: GetBookByID :one
SELECT id, title, published_year, available_quantity, borrow_quantity, created_at, updated_at
FROM "Book"
WHERE id = $1;