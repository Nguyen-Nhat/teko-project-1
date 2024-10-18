package service

import (
	"context"
	"student/internal/repository"
)

type IUniversityService interface {
	CreateUniversity(ctx context.Context) (interface{}, int, error)
}

type universityService struct {
	universityRepository repository.IUniversityRepository
}

func NewUniversityService(universityRepository repository.IUniversityRepository) IUniversityService {
	return &universityService{
		universityRepository: universityRepository,
	}
}
