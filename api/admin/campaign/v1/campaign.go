package v1

import "github.com/gogf/gf/v2/frame/g"

type CampaignItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	SlotId     int64  `json:"slot_id"`
	CreativeId string `json:"creative_id"`
	SiteCode   string `json:"site_code"`
	Priority   int    `json:"priority"`
	Weight     int    `json:"weight"`
	Status     int    `json:"status"`
	StartAt    string `json:"start_at"`
	EndAt      string `json:"end_at"`
	Remark     string `json:"remark"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ListReq struct {
	g.Meta   `path:"/admin/campaigns" method:"get" tags:"Admin/Campaign" summary:"投放列表"`
	Page     int    `json:"page" d:"1"`
	Size     int    `json:"size" d:"20"`
	Keyword  string `json:"keyword"`
	SiteCode string `json:"site_code"`
	SlotId   int64  `json:"slot_id"`
	Status   int    `json:"status" d:"-1"`
}

type ListRes struct {
	List  []CampaignItem `json:"list"`
	Total int            `json:"total"`
}

type CreateReq struct {
	g.Meta     `path:"/admin/campaigns" method:"post" tags:"Admin/Campaign" summary:"创建投放"`
	Name       string `json:"name" v:"required#name必填"`
	SlotId     int64  `json:"slot_id" v:"required|min:1#slot_id必填"`
	CreativeId string `json:"creative_id" v:"required#creative_id必填"`
	SiteCode   string `json:"site_code"`
	Priority   int    `json:"priority" d:"100"`
	Weight     int    `json:"weight" d:"100"`
	Status     int    `json:"status" d:"1"`
	StartAt    string `json:"start_at"`
	EndAt      string `json:"end_at"`
	Remark     string `json:"remark"`
}

type CreateRes struct {
	Id string `json:"id"`
}

type DetailReq struct {
	g.Meta `path:"/admin/campaigns/{id}" method:"get" tags:"Admin/Campaign" summary:"投放详情"`
	Id     string `json:"id" in:"path" v:"required"`
}

type DetailRes struct {
	CampaignItem
}

type UpdateReq struct {
	g.Meta     `path:"/admin/campaigns/{id}" method:"put" tags:"Admin/Campaign" summary:"更新投放"`
	Id         string `json:"id" in:"path" v:"required"`
	Name       string `json:"name" v:"required#name必填"`
	SlotId     int64  `json:"slot_id" v:"required|min:1"`
	CreativeId string `json:"creative_id" v:"required"`
	SiteCode   string `json:"site_code"`
	Priority   int    `json:"priority" d:"100"`
	Weight     int    `json:"weight" d:"100"`
	Status     int    `json:"status" d:"1"`
	StartAt    string `json:"start_at"`
	EndAt      string `json:"end_at"`
	Remark     string `json:"remark"`
}

type UpdateRes struct {
	Id string `json:"id"`
}

type SetStatusReq struct {
	g.Meta `path:"/admin/campaigns/{id}/status" method:"post" tags:"Admin/Campaign" summary:"启停投放"`
	Id     string `json:"id" in:"path" v:"required"`
	Status int    `json:"status" v:"required|in:0,1"`
}

type SetStatusRes struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}
