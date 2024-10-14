//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"library/internal/controller"
	"library/internal/service"
)

func InitGenreRouterHandler() *controller.GenreController {
	wire.Build(
		service.NewGenreService,
		controller.NewGenreController,
	)
	return new(controller.GenreController)
}
