package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"student/internal/dto/res"
	"student/internal/model"
	"student/internal/service"
	"student/mocks"
	"student/pkg/response"
	"testing"
	"time"
)

func TestGetUniversityById(t *testing.T) {
	mockRepo := new(mocks.IUniversityRepository)
	uniService := service.NewUniversityService(mockRepo)
	ctx := context.Background()

	t.Run("WhenUniversityExists_ShouldReturnUniversity", func(t *testing.T) {
		inputUniversityId := 1
		returnUniversity := &model.University{
			ID:                1,
			Name:              "Test University",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			EstablishmentYear: 2024,
		}
		var returnError error = nil
		expectedResult := &res.UniversityDto{
			Id:                1,
			Name:              "Test University",
			EstablishmentYear: 2024,
		}
		expectedStatusCode := response.CodeSuccess

		mockRepo.On("FindByID", ctx, inputUniversityId).Return(returnUniversity, returnError)

		result, statusCode, err := uniService.GetUniversityById(ctx, inputUniversityId)

		require.NoError(t, err)
		assert.Equal(t, statusCode, expectedStatusCode)
		assert.Equal(t, *expectedResult, *result)

		mockRepo.AssertExpectations(t)
	})

	t.Run("WhenUniversityDoesNotExist_ShouldReturnNotFound", func(t *testing.T) {
		// Giả lập không tìm thấy university
		mockRepo.On("FindByID", ctx, 1).Return(nil, nil)

		result, statusCode, err := uniService.GetUniversityById(ctx, 1)

		assert.Nil(t, result)
		assert.Equal(t, response.CodeUniversityNotFound, statusCode)
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("WhenDatabaseError_ShouldReturnError", func(t *testing.T) {
		// Giả lập repository trả về lỗi khi gọi FindByID
		mockRepo.On("FindByID", ctx, 1).Return(nil, errors.New("db error"))

		result, statusCode, err := uniService.GetUniversityById(ctx, 1)

		assert.Nil(t, result)
		assert.Equal(t, response.CodeUniversityNotFound, statusCode)
		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})
}
