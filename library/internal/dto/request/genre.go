package request

type GenrePostDto struct {
	Name string `json:"name" binding:"required,min=1"`
}