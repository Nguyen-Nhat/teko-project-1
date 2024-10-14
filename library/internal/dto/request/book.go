package request

type BookPostDto struct {
	Title         string `json:"title" binding:"required,min=1"`
	PublishedYear int32  `json:"published_year" binding:"required,gte=0"`
}
