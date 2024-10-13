//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"student/internal/controller"
	"student/internal/repository"
	"student/internal/service"
)

func InitUniversityRouterHandler() *controller.UniversityController {
	wire.Build(
		repository.NewUniversityRepository,
		service.NewUniversityService,
		controller.NewUniversityController,
	)
	return new(controller.UniversityController)
}
