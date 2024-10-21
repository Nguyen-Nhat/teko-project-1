package res

import (
	"student/internal/model"
	"time"
)

type StudentDto struct {
	Id             int           `json:"id"`
	FullName       string        `json:"full_name"`
	Sex            int           `json:"sex"`
	Dob            time.Time     `json:"dob"`
	EnrollmentYear int           `json:"enrollment_year"`
	University     UniversityDto `json:"university"`
}

func (s *StudentDto) FromModel(student model.Student) {
	s.Id = student.ID
	s.FullName = student.FullName
	s.Sex = student.Sex
	s.Dob = student.DOB
	s.EnrollmentYear = student.EnrollmentYear
	s.University.FromModel(student.University)
}
