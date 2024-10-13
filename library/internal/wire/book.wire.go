//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"library/internal/controller"
	"library/internal/service"
)

func InitBookRouterHandler() *controller.BookController {
	wire.Build(
		service.NewBookService,
		controller.NewBookController,
	)
	return new(controller.BookController)
}
