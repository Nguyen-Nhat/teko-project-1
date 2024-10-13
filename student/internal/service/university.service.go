package service

import "student/internal/repository"

type IUniversityService interface {
}

type universityService struct {
	universityRepository repository.IUniversityRepository
}

func NewUniversityService(universityRepository repository.IUniversityRepository) IUniversityService {
	return &universityService{
		universityRepository: universityRepository,
	}
}
