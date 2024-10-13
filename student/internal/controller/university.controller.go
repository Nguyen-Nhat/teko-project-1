package controller

import "student/internal/service"

type UniversityController struct {
	universityService service.IUniversityService
}

func NewUniversityController(universityService service.IUniversityService) *UniversityController {
	return &UniversityController{
		universityService: universityService,
	}
}
