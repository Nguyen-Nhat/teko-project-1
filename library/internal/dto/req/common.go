package req

type PageInfo struct {
	Page int `json:"page" form:"page" binding:"omitempty,gte=0" `
	Size int `json:"size" form:"size" binding:"omitempty,gte=1" `
}
