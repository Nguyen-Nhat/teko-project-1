//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"library/internal/controller"
	"library/internal/service"
)

func InitAuthorRouterHandler() *controller.AuthorController {
	wire.Build(
		service.NewAuthorService,
		controller.NewAuthorController,
	)
	return new(controller.AuthorController)
}
