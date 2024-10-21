package req

import "time"

type StudentPostDto struct {
	FullName       string    `json:"full_name" binding:"required,min=1"`
	Sex            int       `json:"sex" binding:"required,gte=0,lte=2"`
	Dob            time.Time `json:"dob" binding:"required"`
	UniversityId   int       `json:"university_id" binding:"required,gte=0"`
	EnrollmentYear int       `json:"enrollment_year" binding:"required,gt=0"`
}

type StudentPageDto struct {
	PageInfo
	UniversityId   int `json:"university_id" form:"university_id" binding:"required,gte=0"`
	EnrollmentYear int `json:"enrollment_year" form:"enrollment_year" binding:"required,gt=0"`
}
