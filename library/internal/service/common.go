package service

import (
	"context"
	"database/sql"
	"library/internal/database"
	"library/pkg/response"
)

func withTransaction[T any](
	ctx context.Context, db *sql.DB, repository *database.Queries,
	fn func(q *database.Queries) (T, int, error)) (T, int, error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		var zero T
		return zero, response.CodeInternalServerError, err
	}
	q := repository.WithTx(tx)

	result, code, err := fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			var zero T
			return zero, response.CodeInternalServerError, rbErr
		}
		return result, code, err
	}
	if err := tx.Commit(); err != nil {
		var zero T
		return zero, response.CodeInternalServerError, err
	}

	return result, code, nil
}
