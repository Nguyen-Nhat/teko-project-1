package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"math"
	"student/global"
	"student/internal/dto/req"
	"student/internal/dto/res"
	"student/internal/model"
	"time"
)

type IStudentRepository interface {
	Create(ctx context.Context, student *model.Student) error
	FindByID(ctx context.Context, id int) (*model.Student, error)
	FindPageByUniIdAndEnrollYear(ctx context.Context, universityId int, enrollYear int, page req.PageInfo) (*res.PageResult[model.Student], error)
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

func (s *studentRepository) FindPageByUniIdAndEnrollYear(ctx context.Context,
	universityId int,
	enrollYear int,
	page req.PageInfo,
) (*res.PageResult[model.Student], error) {

	var students []model.Student
	var total int64

	if err := s.db.Where("university_id = ? AND enrollment_year = ?", universityId, enrollYear).Count(&total).Error; err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(page.Size)))

	offset := (page.Page - 1) * page.Size

	if err := s.db.WithContext(ctx).Preload("University").
		Where("university_id = ? AND enrollment_year = ?", universityId, enrollYear).
		Offset(int(offset)).Limit(int(page.Size)).
		Find(&students).Error; err != nil {
		return nil, err
	}

	result := &res.PageResult[model.Student]{
		List:      students,
		TotalPage: totalPage,
		Page:      int(page.Page),
		Size:      int(page.Size),
	}

	return result, nil
}
