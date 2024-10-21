package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"student/internal/dto/req"
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

	type test struct {
		name               string
		inputUniversityId  int
		returnUniversity   *model.University
		returnError        error
		expectedResult     *res.UniversityDto
		expectedStatusCode int
		expectedError      error
	}
	testcases := []test{
		{
			name:              "WhenUniversityExists_ShouldReturnUniversity",
			inputUniversityId: 1,
			returnUniversity: &model.University{
				ID:                1,
				Name:              "Test University",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
				EstablishmentYear: 2024,
			},
			returnError: nil,
			expectedResult: &res.UniversityDto{
				Id:                1,
				Name:              "Test University",
				EstablishmentYear: 2024,
			},
			expectedStatusCode: response.CodeSuccess,
			expectedError:      nil,
		},
		{
			name:               "WhenUniversityDoesNotExist_ShouldReturnNotFound",
			inputUniversityId:  1,
			returnUniversity:   nil,
			returnError:        nil,
			expectedResult:     nil,
			expectedStatusCode: response.CodeUniversityNotFound,
			expectedError:      nil,
		},
		{
			name:               "WhenDatabaseError_ShouldReturnInternalServerError",
			inputUniversityId:  1,
			returnUniversity:   nil,
			returnError:        errors.New("DB error"),
			expectedResult:     nil,
			expectedStatusCode: response.CodeInternalServerError,
			expectedError:      errors.New("DB error"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.On("FindByID", ctx, tc.inputUniversityId).Return(tc.returnUniversity, tc.returnError)

			result, statusCode, err := uniService.GetUniversityById(ctx, tc.inputUniversityId)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedStatusCode, statusCode)
			assert.Equal(t, tc.expectedError, err)

			mockRepo.AssertExpectations(t)
			mockRepo.ExpectedCalls = nil
		})
	}
}
func TestCreateUniversity(t *testing.T) {
	mockRepo := new(mocks.IUniversityRepository)
	uniService := service.NewUniversityService(mockRepo)
	ctx := context.Background()
	type test struct {
		name               string
		inputRequestBody   *req.UniversityPostDto
		returnError        error
		expectedResult     *res.UniversityDto
		expectedStatusCode int
		expectedError      error
	}
	testcases := []test{
		{
			name: "WhenDatabaseConnected_ShouldReturnUniversity",
			inputRequestBody: &req.UniversityPostDto{
				Name:              "HCMUS",
				EstablishmentYear: 2000,
			},
			returnError: nil,
			expectedResult: &res.UniversityDto{
				Name:              "HCMUS",
				EstablishmentYear: 2000,
			},
			expectedStatusCode: response.CodeCreated,
			expectedError:      nil,
		},
		{
			name: "WhenDatabaseError_ShouldReturnInternalServerError",
			inputRequestBody: &req.UniversityPostDto{
				Name:              "HCMUS",
				EstablishmentYear: 2000,
			},
			returnError:        errors.New("DB error"),
			expectedResult:     nil,
			expectedStatusCode: response.CodeInternalServerError,
			expectedError:      errors.New("DB error"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			data := &model.University{
				Name:              tc.inputRequestBody.Name,
				EstablishmentYear: tc.inputRequestBody.EstablishmentYear,
			}
			mockRepo.On("Create", ctx, data).Return(tc.returnError)
			result, statusCode, err := uniService.CreateUniversity(ctx, tc.inputRequestBody)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedStatusCode, statusCode)
			assert.Equal(t, tc.expectedError, err)

			mockRepo.AssertExpectations(t)
			mockRepo.ExpectedCalls = nil
		})
	}
}
