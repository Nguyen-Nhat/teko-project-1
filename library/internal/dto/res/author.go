package res

import (
	"library/internal/database"
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
	// chưa rõ tại sao gọi hàm Value() của sql.Nulltime mà lại chạy Scan() ???
	if author.Dob.Valid {
		a.Dob = &author.Dob.Time
	} else {
		a.Dob = nil
	}
}
