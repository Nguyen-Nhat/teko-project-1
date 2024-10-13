package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"student/global"
	"student/internal/model"
	"time"
)

type IStudentRepository interface {
	Create(ctx context.Context, student *model.Student) error
	FindByID(ctx context.Context, id int) (*model.Student, error)
	FindAll(ctx context.Context) ([]model.Student, error)
	Update(ctx context.Context, student *model.Student) error
	Delete(ctx context.Context, id int) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository() IStudentRepository {
	return &studentRepository{db: global.Db}
}

func (r *studentRepository) Create(ctx context.Context, student *model.Student) error {
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	if err := r.db.WithContext(ctx).Create(student).Error; err != nil {
		return err
	}
	return nil
}

func (r *studentRepository) FindByID(ctx context.Context, id int) (*model.Student, error) {
	var student model.Student
	if err := r.db.WithContext(ctx).Preload("University").First(&student, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) FindAll(ctx context.Context) ([]model.Student, error) {
	var students []model.Student
	if err := r.db.WithContext(ctx).Preload("University").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) Update(ctx context.Context, student *model.Student) error {
	student.UpdatedAt = time.Now()
	if err := r.db.WithContext(ctx).Save(student).Error; err != nil {
		return err
	}
	return nil
}

func (r *studentRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&model.Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
