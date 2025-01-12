// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: book_genre.sql

package database

import (
	"context"
)

const checkBookGenreExists = `-- name: CheckBookGenreExists :one
SELECT EXISTS(
    SELECT  FROM "Book_Genre" bg
    WHERE bg.book_id = $1 AND bg.genre_id = $2
)
`

type CheckBookGenreExistsParams struct {
	BookID  int32
	GenreID int32
}

func (q *Queries) CheckBookGenreExists(ctx context.Context, arg CheckBookGenreExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkBookGenreExists, arg.BookID, arg.GenreID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const insertBookGenre = `-- name: InsertBookGenre :one
INSERT INTO "Book_Genre" (book_id, genre_id)
VALUES ($1,$2)
RETURNING id, book_id, genre_id
`

type InsertBookGenreParams struct {
	BookID  int32
	GenreID int32
}

func (q *Queries) InsertBookGenre(ctx context.Context, arg InsertBookGenreParams) (BookGenre, error) {
	row := q.db.QueryRowContext(ctx, insertBookGenre, arg.BookID, arg.GenreID)
	var i BookGenre
	err := row.Scan(&i.ID, &i.BookID, &i.GenreID)
	return i, err
}

const removeBookGenre = `-- name: RemoveBookGenre :exec
DELETE FROM "Book_Genre" bg
WHERE bg.book_id = $1 AND bg.genre_id = $2
`

type RemoveBookGenreParams struct {
	BookID  int32
	GenreID int32
}

func (q *Queries) RemoveBookGenre(ctx context.Context, arg RemoveBookGenreParams) error {
	_, err := q.db.ExecContext(ctx, removeBookGenre, arg.BookID, arg.GenreID)
	return err
}
