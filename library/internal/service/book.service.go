package service

import (
	"context"
	"library/global"
	"library/internal/database"
)

type IBookService interface {
	GetAll(ctx context.Context) (int, error)
}

type bookService struct {
	repository *database.Queries
}

func NewBookService() IBookService {
	return &bookService{
		repository: database.New(global.Db),
	}
}
func (bs *bookService) GetAll(ctx context.Context) (int, error) {
	//result, err := bs.repository.GetAllAuthors(ctx)

	return 1, nil
}
