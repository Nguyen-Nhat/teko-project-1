package res

type PageResult struct {
	List      interface{} `json:"list"`
	TotalPage int         `json:"total_page"`
	Page      int         `json:"page"`
	Size      int         `json:"size"`
}
