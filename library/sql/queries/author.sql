-- name: CreateAuthor :one
INSERT INTO "Author" (fullname, dob)
VALUES ($1,$2)
RETURNING *;

-- name: GetAuthorByID :one
SELECT a.*
FROM "Author" as a
WHERE id = $1;

-- name: GetAllAuthors :many
SELECT a.*
FROM "Author" as a;

-- name: GetPageAuthors :many
SELECT a.*
FROM "Author" as a
LIMIT sqlc.arg(size)
OFFSET sqlc.arg(size) * (sqlc.arg(page)::INTEGER - 1);

-- name: GetAuthorsByBookID :many
SELECT a.*
FROM "Author" as a INNER JOIN "Book_Author" as ba
ON a.id = ba.author_id;