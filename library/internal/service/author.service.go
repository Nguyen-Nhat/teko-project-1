package service

import (
	"context"
	"database/sql"
	"library/global"
	"library/internal/database"
	"library/internal/dto/req"
	"library/pkg/response"
)

type IAuthorService interface {
	GetAuthorById(ctx context.Context, id int) (*database.Author, int, error)
	CreateAuthor(ctx context.Context, data *req.AuthorPostDto) (*database.Author, int, error)
}

type authorService struct {
	db         *sql.DB
	repository *database.Queries
}

func NewAuthorService() IAuthorService {
	return &authorService{
		db:         global.Db,
		repository: database.New(global.Db),
	}
}

func (as *authorService) CreateAuthor(ctx context.Context, data *req.AuthorPostDto) (*database.Author, int, error) {
	tx, err := as.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, response.CodeInternalServerError, err
	}
	q := as.repository.WithTx(tx)

	dob := sql.NullTime{}
	_ = dob.Scan(data.Dob)

	params := database.CreateAuthorParams{
		Fullname: data.FullName,
		Dob:      dob,
	}
	result, err := q.CreateAuthor(ctx, params)
	if err != nil {
		_ = tx.Rollback()
		return nil, response.CodeCannotCreateAuthor, err
	}
	_ = tx.Commit()
	return &result, response.CodeCreated, nil
}
func (as *authorService) GetAuthorById(ctx context.Context, id int) (*database.Author, int, error) {
	result, err := as.repository.GetAuthorByID(ctx, int32(id))
	if err != nil {
		return nil, response.CodeGenreNotFound, err
	}
	return &result, response.CodeSuccess, nil
}
