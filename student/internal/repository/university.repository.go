package repository

type IUniversityRepository interface {
}

type universityRepository struct{}

func NewUniversityRepository() IUniversityRepository {
	return &universityRepository{}
}
