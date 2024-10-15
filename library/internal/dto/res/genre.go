package res

import "library/internal/database"

type GenreDto struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func (g *GenreDto) FromModel(genre *database.Genre) {
	g.Id = genre.ID
	g.Name = genre.Name
}
