package res

type PageResult[T any] struct {
	List      []T `json:"list"`
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
	Size      int `json:"size"`
}
