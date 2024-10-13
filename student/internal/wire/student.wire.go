//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"student/internal/controller"
	"student/internal/repository"
	"student/internal/service"
)

func InitStudentRouterHandler() *controller.StudentController {
	wire.Build(
		repository.NewStudentRepository,
		service.NewStudentService,
		controller.NewStudentController,
	)
	return new(controller.StudentController)
}
