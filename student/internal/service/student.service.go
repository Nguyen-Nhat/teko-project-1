package service

import (
	"context"
	"student/internal/repository"
)

type IStudentService interface {
	CreateStudent(ctx context.Context) (interface{}, int, error)
	GetStudentById(ctx context.Context, studentId int32) (interface{}, int, error)
}

type studentService struct {
	studentRepository repository.IStudentRepository
}

func NewStudentService(studentRepository repository.IStudentRepository) IStudentService {
	return &studentService{
		studentRepository: studentRepository,
	}
}
