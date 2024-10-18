package res

import (
	"library/internal/database"
	"library/internal/util"
	"time"
)

type AuthorDto struct {
	ID       int32      `json:"id"`
	FullName string     `json:"full_name"`
	Dob      *time.Time `json:"dob"`
}

func (a *AuthorDto) FromModel(author database.Author) {
	a.ID = author.ID
	a.FullName = author.Fullname
	a.Dob = util.FromNullTime(author.Dob)
}
