package router

type RouterGroup struct {
	BookRouter
	GenreRouter
	AuthorRouter
	BorrowBookRouter
}

var RouterGroupApp = new(RouterGroup)
