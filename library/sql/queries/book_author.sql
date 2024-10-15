-- name: CheckBookAuthorExists :one
SELECT EXISTS(
    SELECT ba.* FROM "Book_Author" ba
    WHERE ba.book_id = $1 AND ba.author_id = $2
);
-- name: InsertBookAuthor :one
INSERT INTO "Book_Author" (book_id, author_id)
VALUES ($1,$2)
RETURNING *;

-- name: RemoveBookAuthor :exec
DELETE FROM "Book_Author" ba
WHERE ba.book_id = $1 AND ba.author_id = $2;