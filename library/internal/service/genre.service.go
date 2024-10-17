package service

import (
	"context"
	"database/sql"
	"library/global"
	"library/internal/database"
	"library/internal/dto/req"
	"library/pkg/response"
)

type IGenreService interface {
	GetGenreById(ctx context.Context, id int) (*database.Genre, int, error)
	CreateGenre(ctx context.Context, data *req.GenrePostDto) (*database.Genre, int, error)
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

func (gs *genreService) CreateGenre(ctx context.Context, data *req.GenrePostDto) (*database.Genre, int, error) {
	return withTransaction(ctx, gs.db, gs.repository, func(q *database.Queries) (*database.Genre, int, error) {
		result, err := q.CreateGenre(ctx, data.Name)
		if err != nil {
			return nil, response.CodeCannotCreateGenre, err
		}
		return &result, response.CodeSuccess, nil
	})
}

func (gs *genreService) GetGenreById(ctx context.Context, id int) (*database.Genre, int, error) {
	result, err := gs.repository.GetGenreByID(ctx, int32(id))
	if err != nil {
		return nil, response.CodeGenreNotFound, err
	}
	return &result, response.CodeSuccess, nil
}
