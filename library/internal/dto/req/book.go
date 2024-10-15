package req

type BookPostDto struct {
	Title         string `json:"title" binding:"required,min=1"`
	PublishedYear int32  `json:"published_year" binding:"required,gte=0"`
}

type BookPageDto struct {
	Title    *string `json:"title" form:"title" binding:"omitempty"`
	GenreId  *int32  `json:"genre_id" form:"genre_id" binding:"omitempty,gte=0"`
	AuthorId *int32  `json:"author_id" form:"author_id" binding:"omitempty,gte=0"`
	PageInfo
}
