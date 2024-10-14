package request

import "time"

type AuthorPostDto struct {
	Fullname string     `json:"full_name" binding:"required,min=1"`
	Dob      *time.Time `json:"dob" binding:"omitempty"`
}
