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

func TestGetStudentById(t *testing.T) {
	mockStudentRepo := new(mocks.IStudentRepository)
	mockUniRepo := new(mocks.IUniversityRepository)
	studentService := service.NewStudentService(mockStudentRepo, mockUniRepo)
	ctx := context.Background()
	type test struct {
		name               string
		inputStudentId     int
		returnStudent      *model.Student
		returnError        error
		expectedResult     *res.StudentDto
		expectedStatusCode int
		expectedError      error
	}
	now := time.Now()
	testcases := []test{
		{
			name:           "WhenStudentExists_ShouldReturnStudent",
			inputStudentId: 1,
			returnStudent: &model.Student{
				ID:             1,
				DOB:            now,
				EnrollmentYear: 2021,
				Sex:            1,
				FullName:       "Nguyễn Nhật",
				CreatedAt:      now,
				UpdatedAt:      now,
				University: model.University{
					ID:                1,
					Name:              "HCMUS",
					EstablishmentYear: 1999,
					CreatedAt:         now,
					UpdatedAt:         now,
				},
			},
			returnError: nil,
			expectedResult: &res.StudentDto{
				Id:             1,
				Dob:            now,
				EnrollmentYear: 2021,
				Sex:            1,
				FullName:       "Nguyễn Nhật",
				University: res.UniversityDto{
					Id:                1,
					Name:              "HCMUS",
					EstablishmentYear: 1999,
				},
			},
			expectedStatusCode: response.CodeSuccess,
			expectedError:      nil,
		},
		{
			name:               "WhenStudentDoesNotExist_ShouldReturnNotFound",
			inputStudentId:     1,
			returnStudent:      nil,
			returnError:        nil,
			expectedResult:     nil,
			expectedStatusCode: response.CodeStudentNotFound,
			expectedError:      nil,
		},
		{
			name:               "WhenDatabaseError_ShouldReturnInternalServerError",
			inputStudentId:     1,
			returnStudent:      nil,
			returnError:        errors.New("DB error"),
			expectedResult:     nil,
			expectedStatusCode: response.CodeInternalServerError,
			expectedError:      errors.New("DB error"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStudentRepo.On("FindByID", ctx, tc.inputStudentId).Return(tc.returnStudent, tc.returnError)

			result, statusCode, err := studentService.GetStudentById(ctx, tc.inputStudentId)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedStatusCode, statusCode)
			assert.Equal(t, tc.expectedError, err)

			mockStudentRepo.AssertExpectations(t)
			mockStudentRepo.ExpectedCalls = nil
		})
	}
}

func TestCreateStudent(t *testing.T) {
	mockStudentRepo := new(mocks.IStudentRepository)
	mockUniRepo := new(mocks.IUniversityRepository)
	studentService := service.NewStudentService(mockStudentRepo, mockUniRepo)
	ctx := context.Background()
	type test struct {
		name               string
		inputRequestBody   *req.StudentPostDto
		returnUni          *model.University
		returnStudentError error
		returnUniError     error
		expectedResult     *res.StudentDto
		expectedStatusCode int
		expectedError      error
	}
	var sex = 1
	var now = time.Now()
	testcases := []test{
		{
			name: "WhenAllCriteriaSatisfaction_ShouldReturnStudent",
			inputRequestBody: &req.StudentPostDto{
				FullName:       "Nguyễn Nhật",
				Sex:            &sex,
				Dob:            now,
				UniversityId:   1,
				EnrollmentYear: 2021,
			},
			returnUni: &model.University{
				ID:                1,
				Name:              "HCMUS",
				EstablishmentYear: 1999,
				CreatedAt:         now,
				UpdatedAt:         now,
			},
			returnStudentError: nil,
			returnUniError:     nil,
			expectedError:      nil,
			expectedStatusCode: response.CodeSuccess,
			expectedResult: &res.StudentDto{
				FullName:       "Nguyễn Nhật",
				Sex:            sex,
				Dob:            now,
				EnrollmentYear: 2021,
				University: res.UniversityDto{
					Id:                1,
					Name:              "HCMUS",
					EstablishmentYear: 1999,
				},
			},
		},
		{
			name: "WhenUniversityDoesNotExist_ShouldReturnNotFound",
			inputRequestBody: &req.StudentPostDto{
				FullName:       "Nguyễn Nhật",
				Sex:            &sex,
				Dob:            now,
				UniversityId:   1,
				EnrollmentYear: 2021,
			},
			returnUni:          nil,
			returnStudentError: nil,
			returnUniError:     nil,
			expectedError:      nil,
			expectedStatusCode: response.CodeUniversityNotFound,
			expectedResult:     nil,
		},
		{
			name: "WhenDatabaseErrorAtGetUni_ShouldReturnInternalServerError",
			inputRequestBody: &req.StudentPostDto{
				FullName:       "Nguyễn Nhật",
				Sex:            &sex,
				Dob:            now,
				UniversityId:   1,
				EnrollmentYear: 2021,
			},
			returnUni:          nil,
			returnStudentError: errors.New("DB error"),
			returnUniError:     errors.New("DB error"),
			expectedError:      errors.New("DB error"),
			expectedStatusCode: response.CodeInternalServerError,
			expectedResult:     nil,
		},
		{
			name: "WhenDatabaseErrorAtCreateStudent_ShouldReturnInternalServerError",
			inputRequestBody: &req.StudentPostDto{
				FullName:       "Nguyễn Nhật",
				Sex:            &sex,
				Dob:            now,
				UniversityId:   1,
				EnrollmentYear: 2021,
			},
			returnUni: &model.University{
				ID:                1,
				Name:              "HCMUS",
				EstablishmentYear: 1999,
				CreatedAt:         now,
				UpdatedAt:         now,
			},
			returnStudentError: errors.New("DB error"),
			returnUniError:     nil,
			expectedError:      errors.New("DB error"),
			expectedStatusCode: response.CodeInternalServerError,
			expectedResult:     nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockUniRepo.On("FindByID", ctx, tc.inputRequestBody.UniversityId).Return(tc.returnUni, tc.returnUniError)

			if tc.returnUni != nil {
				data := model.Student{
					FullName:       tc.inputRequestBody.FullName,
					Sex:            *tc.inputRequestBody.Sex,
					DOB:            tc.inputRequestBody.Dob,
					EnrollmentYear: tc.inputRequestBody.EnrollmentYear,
					University:     *tc.returnUni,
				}
				mockStudentRepo.On("Create", ctx, &data).Return(tc.returnStudentError)
			}
			result, statusCode, err := studentService.CreateStudent(ctx, tc.inputRequestBody)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedStatusCode, statusCode)
			assert.Equal(t, tc.expectedError, err)

			mockStudentRepo.AssertExpectations(t)
			mockUniRepo.AssertExpectations(t)
			mockStudentRepo.ExpectedCalls = nil
			mockUniRepo.ExpectedCalls = nil
		})
	}
}

func TestGetPageStudentWithFilter(t *testing.T) {
	mockStudentRepo := new(mocks.IStudentRepository)
	mockUniRepo := new(mocks.IUniversityRepository)
	studentService := service.NewStudentService(mockStudentRepo, mockUniRepo)
	ctx := context.Background()
	type test struct {
		name               string
		inputRequestQuery  *req.StudentPageDto
		returnPageStudent  *res.PageResult[model.Student]
		returnError        error
		expectedResult     *res.PageResult[res.StudentDto]
		expectedStatusCode int
		expectedError      error
	}
	now := time.Now()
	testcases := []test{
		{
			name: "WhenDatabaseConnected_ShouldReturnPageStudent",
			inputRequestQuery: &req.StudentPageDto{
				UniversityId:   1,
				EnrollmentYear: 2021,
				PageInfo: req.PageInfo{
					Page: 1,
					Size: 2,
				},
			},
			returnPageStudent: &res.PageResult[model.Student]{
				TotalPage: 2,
				Page:      1,
				Size:      2,
				List: []model.Student{
					{
						ID:       1,
						FullName: "Nguyễn Văn A",
						DOB:      now,
						Sex:      1,
						University: model.University{
							ID:                1,
							Name:              "HCMUS",
							EstablishmentYear: 1999,
							UpdatedAt:         now,
							CreatedAt:         now,
						},
						EnrollmentYear: 2021,
						UpdatedAt:      now,
						CreatedAt:      now,
					},
					{
						ID:       2,
						FullName: "Nguyễn Thị B",
						DOB:      now,
						Sex:      0,
						University: model.University{
							ID:                1,
							Name:              "HCMUS",
							EstablishmentYear: 1999,
							UpdatedAt:         now,
							CreatedAt:         now,
						},
						EnrollmentYear: 2021,
						UpdatedAt:      now,
						CreatedAt:      now,
					},
				},
			},
			returnError: nil,
			expectedResult: &res.PageResult[res.StudentDto]{
				TotalPage: 2,
				Page:      1,
				Size:      2,
				List: []res.StudentDto{
					{
						Id:       1,
						FullName: "Nguyễn Văn A",
						Dob:      now,
						Sex:      1,
						University: res.UniversityDto{
							Id:                1,
							Name:              "HCMUS",
							EstablishmentYear: 1999,
						},
						EnrollmentYear: 2021,
					},
					{
						Id:       2,
						FullName: "Nguyễn Thị B",
						Dob:      now,
						Sex:      0,
						University: res.UniversityDto{
							Id:                1,
							Name:              "HCMUS",
							EstablishmentYear: 1999,
						},
						EnrollmentYear: 2021,
					},
				},
			},
			expectedError:      nil,
			expectedStatusCode: response.CodeSuccess,
		},
		{
			name: "WhenDatabaseError_ShouldReturnInternalServerError",
			inputRequestQuery: &req.StudentPageDto{
				UniversityId:   1,
				EnrollmentYear: 2021,
				PageInfo: req.PageInfo{
					Page: 1,
					Size: 3,
				},
			},
			returnPageStudent:  nil,
			returnError:        errors.New("DB error"),
			expectedResult:     nil,
			expectedStatusCode: response.CodeInternalServerError,
			expectedError:      errors.New("DB error"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStudentRepo.On("FindPageByUniIdAndEnrollYear",
				ctx,
				tc.inputRequestQuery.UniversityId,
				tc.inputRequestQuery.EnrollmentYear,
				tc.inputRequestQuery.PageInfo,
			).Return(tc.returnPageStudent, tc.returnError)
			result, statusCode, err := studentService.GetPageStudentWithFilter(ctx, tc.inputRequestQuery)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedStatusCode, statusCode)
			assert.Equal(t, tc.expectedError, err)

			mockStudentRepo.AssertExpectations(t)
			mockStudentRepo.ExpectedCalls = nil
		})
	}
}
