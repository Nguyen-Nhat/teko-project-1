//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"library/internal/controller"
	"library/internal/service"
)

func InitBorrowBookRouterHandler() *controller.BorrowBookController {
	wire.Build(
		service.NewBorrowBookService,
		controller.NewBorrowBookController,
	)
	return new(controller.BorrowBookController)
}
