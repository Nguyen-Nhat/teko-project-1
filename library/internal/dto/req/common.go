package req

import "library/global"

type PageInfo struct {
	Page int `json:"page" form:"page" binding:"min=0"`
	Size int `json:"size" form:"size" binding:"min=1"`
}

func (p *PageInfo) SetDefaultPageInfo() {
	p.Page = global.Config.Server.Pageable.DefaultPage
	p.Size = global.Config.Server.Pageable.DefaultSize
}
