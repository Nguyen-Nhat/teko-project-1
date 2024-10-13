package router

type RouterGroup struct {
	BookRouter
	GenreRouter
	AuthorRouter
}

var RouterGroupApp = new(RouterGroup)
