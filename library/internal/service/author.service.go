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

type IAuthorService interface {
	GetAuthorById(ctx context.Context, id int) (*res.AuthorDto, int, error)
	CreateAuthor(ctx context.Context, data *req.AuthorPostDto) (*res.AuthorDto, int, error)
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

func (as *authorService) CreateAuthor(ctx context.Context, data *req.AuthorPostDto) (*res.AuthorDto, int, error) {
	return withTransaction(ctx, as.db, as.repository, func(q *database.Queries) (*res.AuthorDto, int, error) {
		dob := sql.NullTime{}
		err := dob.Scan(data.Dob)
		if err != nil {
			return nil, response.CodeInternalServerError, err
		}

		params := database.CreateAuthorParams{
			Fullname: data.FullName,
			Dob:      dob,
		}
		author, err := q.CreateAuthor(ctx, params)
		if err != nil {
			return nil, response.CodeCannotCreateAuthor, err
		}
		result := res.AuthorDto{}
		result.FromModel(author)
		return &result, response.CodeCreated, nil
	})
}
func (as *authorService) GetAuthorById(ctx context.Context, id int) (*res.AuthorDto, int, error) {
	author, err := as.repository.GetAuthorByID(ctx, int32(id))
	if err != nil {
		return nil, response.CodeGenreNotFound, err
	}
	result := res.AuthorDto{}
	result.FromModel(author)
	return &result, response.CodeSuccess, nil
}
