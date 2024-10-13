package controller

import "student/internal/service"

type StudentController struct {
	studentService service.IStudentService
}

func NewStudentController(studentService service.IStudentService) *StudentController {
	return &StudentController{
		studentService: studentService,
	}
}
