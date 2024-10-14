-- name: CreateBook :one
INSERT INTO "Book" (title, published_year)
VALUES ($1,$2)
RETURNING *;

-- name: GetAllBooks :many
SELECT b.*
FROM "Book" as b;

-- name: GetBookByID :one
SELECT b.*
FROM "Book" as b
WHERE id = $1;

-- name: GetPageBooks :many
SELECT b.*
FROM "Book" as b
LIMIT sqlc.arg(size)
OFFSET sqlc.arg(size) * (sqlc.arg(page)::INTEGER - 1);
