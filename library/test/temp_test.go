package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"library/internal/dto/req"
	"library/mocks"
	"testing"
)

func TestCreateGenre(t *testing.T) {
	mockGenreService := new(mocks.IGenreService)
	ctx := context.Background()

	inputData := &req.GenrePostDto{Name: ""}

	mockGenreService.On("CreateGenre", ctx, inputData).Return(nil, 400, errors.New(""))

	_, statusCode, err := mockGenreService.CreateGenre(ctx, inputData)

	assert.Equal(t, 400, statusCode)
	assert.EqualError(t, err, "")

	mockGenreService.AssertExpectations(t)
}
