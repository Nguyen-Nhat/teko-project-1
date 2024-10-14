package service

import (
	"context"
	"database/sql"
	"library/global"
	"library/internal/database"
	"library/internal/dto/request"
	"library/pkg/response"
)

type IGenreService interface {
	CreateGenre(ctx context.Context, data *request.GenrePostDto) (*database.Genre, int, error)
}

type genreService struct {
	db         *sql.DB
	repository *database.Queries
}

func NewGenreService() IGenreService {
	return &genreService{
		db:         global.Db,
		repository: database.New(global.Db),
	}
}

func (gs *genreService) CreateGenre(ctx context.Context, data *request.GenrePostDto) (*database.Genre, int, error) {
	tx, _ := gs.db.BeginTx(ctx, nil)
	q := gs.repository.WithTx(tx)
	result, err := q.CreateGenre(ctx, data.Name)
	if err != nil {
		_ = tx.Rollback()
		return nil, response.CodeCannotCreateGenre, err
	}
	_ = tx.Commit()
	return &result, response.CodeCreated, nil
}
