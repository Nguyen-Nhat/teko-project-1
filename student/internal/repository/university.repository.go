package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"student/global"
	"student/internal/model"
)

type IUniversityRepository interface {
	Create(ctx context.Context, university *model.University) error
	FindByID(ctx context.Context, id int) (*model.University, error)
}

type universityRepository struct {
	db *gorm.DB
}

func NewUniversityRepository() IUniversityRepository {
	return &universityRepository{db: global.Db}
}
func (r *universityRepository) Create(ctx context.Context, university *model.University) error {
	if err := r.db.WithContext(ctx).Create(university).Error; err != nil {
		return err
	}
	return nil
}

func (r *universityRepository) FindByID(ctx context.Context, id int) (*model.University, error) {
	var university model.University
	if err := r.db.WithContext(ctx).First(&university, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &university, nil
}
