package util

import (
	"database/sql"
	"time"
)

func ToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{
			Valid: false,
		}
	}
	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}
func ToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}
func ToNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{
			Int32: 0,
			Valid: false,
		}
	}
	return sql.NullInt32{
		Int32: *i,
		Valid: true,
	}
}
