package req

import "time"

type StudentPostDto struct {
	FullName       string    `json:"full_name" binding:"required,min=1"`
	Sex            *int      `json:"sex" binding:"required,min=0,max=2"`
	Dob            time.Time `json:"dob" form:"dob" binding:"required"`
	UniversityId   int       `json:"university_id" binding:"required,min=1"`
	EnrollmentYear int       `json:"enrollment_year" binding:"required,min=1"`
}

type StudentPageDto struct {
	PageInfo
	UniversityId   int `json:"university_id" form:"university_id" binding:"required,min=1"`
	EnrollmentYear int `json:"enrollment_year" form:"enrollment_year" binding:"required,min=1"`
}
