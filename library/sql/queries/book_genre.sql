-- name: CheckBookGenreExists :one
SELECT EXISTS(
    SELECT ba.* FROM "Book_Genre" bg
    WHERE bg.book_id = $1 AND bg.genre_id = $2
);

-- name: InsertBookGenre :one
INSERT INTO "Book_Genre" (book_id, genre_id)
VALUES ($1,$2)
RETURNING *;

-- name: RemoveBookGenre :exec
DELETE FROM "Book_Genre" bg
WHERE bg.book_id = $1 AND bg.genre_id = $2;
