package service

import (
	"context"
	"student/internal/dto/req"
	"student/internal/dto/res"
	"student/internal/model"
	"student/internal/repository"
	"student/pkg/response"
)

type IStudentService interface {
	CreateStudent(ctx context.Context, data *req.StudentPostDto) (*res.StudentDto, int, error)
	GetStudentById(ctx context.Context, studentId int) (*res.StudentDto, int, error)
	GetPageStudentWithFilter(ctx context.Context, data *req.StudentPageDto) (*res.PageResult[res.StudentDto], int, error)
}

type studentService struct {
	studentRepository    repository.IStudentRepository
	universityRepository repository.IUniversityRepository
}

func NewStudentService(studentRepository repository.IStudentRepository, universityRepository repository.IUniversityRepository) IStudentService {
	return &studentService{
		studentRepository:    studentRepository,
		universityRepository: universityRepository,
	}
}
func (ss *studentService) CreateStudent(ctx context.Context, data *req.StudentPostDto) (*res.StudentDto, int, error) {
	_, err := ss.universityRepository.FindByID(ctx, data.UniversityId)
	if err != nil {
		return nil, response.CodeUniversityNotFound, err
	}

	student := model.Student{
		FullName:     data.FullName,
		Sex:          data.Sex,
		DOB:          data.Dob,
		UniversityID: data.UniversityId,
	}
	if err := ss.studentRepository.Create(ctx, &student); err != nil {
		return nil, response.CodeInternalServerError, err
	}
	result := &res.StudentDto{}
	result.FromModel(student)
	return result, response.CodeSuccess, nil
}
func (ss *studentService) GetStudentById(ctx context.Context, studentId int) (*res.StudentDto, int, error) {
	student, err := ss.studentRepository.FindByID(ctx, studentId)
	if err != nil {
		return nil, response.CodeStudentNotFound, err
	}
	result := &res.StudentDto{}
	result.FromModel(*student)
	return result, response.CodeSuccess, nil
}

func (ss *studentService) GetPageStudentWithFilter(ctx context.Context, data *req.StudentPageDto) (*res.PageResult[res.StudentDto], int, error) {
	page, err := ss.studentRepository.FindPageByUniIdAndEnrollYear(ctx, data.UniversityId, data.EnrollmentYear, data.PageInfo)
	if err != nil {
		return nil, response.CodeInternalServerError, err
	}
	students := make([]res.StudentDto, len(page.List))
	for i, student := range page.List {
		students[i].FromModel(student)
	}
	result := &res.PageResult[res.StudentDto]{
		List:      students,
		TotalPage: page.TotalPage,
		Page:      page.Page,
		Size:      page.Size,
	}
	return result, response.CodeSuccess, nil
}
