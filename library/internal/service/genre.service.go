package service

import (
	"context"
	"database/sql"
	"library/global"
	"library/internal/database"
	"library/internal/dto/req"
	"library/internal/dto/res"
	"library/pkg/response"
)

type IGenreService interface {
	GetGenreById(ctx context.Context, id int) (*res.GenreDto, int, error)
	CreateGenre(ctx context.Context, data *req.GenrePostDto) (*res.GenreDto, int, error)
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

func (gs *genreService) CreateGenre(ctx context.Context, data *req.GenrePostDto) (*res.GenreDto, int, error) {
	return withTransaction(ctx, gs.db, gs.repository, func(q *database.Queries) (*res.GenreDto, int, error) {
		genre, err := q.CreateGenre(ctx, data.Name)
		if err != nil {
			return nil, response.CodeCannotCreateGenre, err
		}
		result := res.GenreDto{}
		result.FromModel(genre)
		return &result, response.CodeCreated, nil
	})
}

func (gs *genreService) GetGenreById(ctx context.Context, id int) (*res.GenreDto, int, error) {
	genre, err := gs.repository.GetGenreByID(ctx, int32(id))
	if err != nil {
		return nil, response.CodeGenreNotFound, err
	}
	result := res.GenreDto{}
	result.FromModel(genre)
	return &result, response.CodeSuccess, nil
}
