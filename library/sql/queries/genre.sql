-- name: CreateGenre :one
INSERT INTO "Genre" (name)
VALUES ($1)
RETURNING *;

-- name: GetGenreByID :one
SELECT g.*
FROM "Genre" as g
WHERE id = $1;

-- name: GetAllGenres :many
SELECT g.*
FROM "Genre" as g;

-- name: GetPageGenres :many
SELECT g.*
FROM "Genre" as g
LIMIT sqlc.arg(size)
OFFSET sqlc.arg(size) * (sqlc.arg(page)::INTEGER - 1);

-- name: GetGenreByBookID :many
SELECT g.*
FROM "Genre" as g INNER JOIN "Book_Genre" as bg
ON g.id = bg.genre_id
WHERE bg.book_id = $1;
