package service

import "student/internal/repository"

type IStudentService interface {
}

type studentService struct {
	studentRepository repository.IStudentRepository
}

func NewStudentService(studentRepository repository.IStudentRepository) IStudentService {
	return &studentService{
		studentRepository: studentRepository,
	}
}
