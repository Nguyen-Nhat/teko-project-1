// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"library/internal/controller"
	"library/internal/service"
)

// Injectors from book.wire.go:

func InitBookRouterHandler() *controller.BookController {
	iBookService := service.NewBookService()
	bookController := controller.NewBookController(iBookService)
	return bookController
}
