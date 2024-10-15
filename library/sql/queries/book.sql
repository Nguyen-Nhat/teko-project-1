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
SELECT DISTINCT b.*
FROM "Book" as b
LEFT JOIN "Book_Genre" as bg
ON b.id == bg.book_id
LEFT JOIN "Book_Author" as ba
ON b.id == ba.book_id
WHERE
    (sqlc.narg(title)::text IS NULL OR b.title ILIKE '%' || sqlc.narg(title)::text || '%')
    AND (sqlc.narg(authorId)::int IS NULL OR ba.author_id = sqlc.narg(authorId)::int)
    AND (sqlc.narg(genreId)::int IS NULL OR bg.genre_id = sqlc.narg(genreId)::int)
LIMIT sqlc.arg(size)
OFFSET sqlc.arg(size) * (sqlc.arg(page)::INTEGER - 1);
