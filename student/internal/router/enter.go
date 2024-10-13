package router

type RouterGroup struct {
	StudentRouter
	UniversityRouter
}

var RouterGroupApp = new(RouterGroup)
