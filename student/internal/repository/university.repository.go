package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"student/global"
	"student/internal/model"
	"time"
)

type IUniversityRepository interface {
	Create(ctx context.Context, university *model.University) error
	FindByID(ctx context.Context, id int) (*model.University, error)
	FindAll(ctx context.Context) ([]model.University, error)
	Update(ctx context.Context, university *model.University) error
	Delete(ctx context.Context, id int) error
}

type universityRepository struct {
	db *gorm.DB
}

func NewUniversityRepository() IUniversityRepository {
	return &universityRepository{db: global.Db}
}
func (r *universityRepository) Create(ctx context.Context, university *model.University) error {
	university.CreatedAt = time.Now()
	university.UpdatedAt = time.Now()
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

func (r *universityRepository) FindAll(ctx context.Context) ([]model.University, error) {
	var universities []model.University
	if err := r.db.WithContext(ctx).Find(&universities).Error; err != nil {
		return nil, err
	}
	return universities, nil
}

func (r *universityRepository) Update(ctx context.Context, university *model.University) error {
	university.UpdatedAt = time.Now()
	if err := r.db.WithContext(ctx).Save(university).Error; err != nil {
		return err
	}
	return nil
}

func (r *universityRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&model.University{}, id).Error; err != nil {
		return err
	}
	return nil
}
