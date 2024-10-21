package service

import (
	"context"
	"student/internal/dto/req"
	"student/internal/dto/res"
	"student/internal/model"
	"student/internal/repository"
	"student/pkg/response"
)

type IUniversityService interface {
	CreateUniversity(ctx context.Context, data *req.UniversityPostDto) (*res.UniversityDto, int, error)
	GetUniversityById(ctx context.Context, id int) (*res.UniversityDto, int, error)
}

type universityService struct {
	universityRepository repository.IUniversityRepository
}

func NewUniversityService(universityRepository repository.IUniversityRepository) IUniversityService {
	return &universityService{
		universityRepository: universityRepository,
	}
}

func (us *universityService) CreateUniversity(ctx context.Context, data *req.UniversityPostDto) (*res.UniversityDto, int, error) {
	university := model.University{
		Name:              data.Name,
		EstablishmentYear: data.EstablishmentYear,
	}
	err := us.universityRepository.Create(ctx, &university)
	if err != nil {
		return nil, response.CodeInternalServerError, err
	}
	result := &res.UniversityDto{}
	result.FromModel(university)
	return result, response.CodeCreated, nil
}
func (us *universityService) GetUniversityById(ctx context.Context, id int) (*res.UniversityDto, int, error) {
	university, err := us.universityRepository.FindByID(ctx, id)
	if err != nil {
		return nil, response.CodeInternalServerError, err
	}
	if university == nil {
		return nil, response.CodeUniversityNotFound, err
	}
	result := &res.UniversityDto{}
	result.FromModel(*university)
	return result, response.CodeSuccess, nil
}
