package response

import "database/sql"

type AuthorDto struct {
	ID       int32        `json:"id"`
	FullName string       `json:"full_name"`
	Dob      sql.NullTime `json:"dob"`
}
