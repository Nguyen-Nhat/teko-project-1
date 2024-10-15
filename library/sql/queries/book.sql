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


-- name: GetPageBookWithFilter :many
SELECT b.*
FROM "Book" as b
INNER JOIN "Book_Genre" as bg
ON b.id == bg.book_id
INNER JOIN "Book_Author" as ba
ON b.id == ba.book_id
LIMIT sqlc.arg(size)
OFFSET sqlc.arg(size) * (sqlc.arg(page)::INTEGER - 1);
