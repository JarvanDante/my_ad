package v1

import "github.com/gogf/gf/v2/frame/g"

type SlotItem struct {
	Id        int64  `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	SlotType  string `json:"slot_type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Status    int    `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListReq struct {
	g.Meta  `path:"/admin/slots" method:"get" tags:"Admin/Slot" summary:"广告位列表"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	Keyword string `json:"keyword"`
	Status  int    `json:"status" d:"-1"`
}

type ListRes struct {
	List  []SlotItem `json:"list"`
	Total int        `json:"total"`
}

type CreateReq struct {
	g.Meta   `path:"/admin/slots" method:"post" tags:"Admin/Slot" summary:"创建广告位"`
	Code     string `json:"code" v:"required#code必填"`
	Name     string `json:"name" v:"required#name必填"`
	SlotType string `json:"slot_type" d:"banner"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Status   int    `json:"status" d:"1"`
	Remark   string `json:"remark"`
}

type CreateRes struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type DetailReq struct {
	g.Meta `path:"/admin/slots/{id}" method:"get" tags:"Admin/Slot" summary:"广告位详情"`
	Id     int64 `json:"id" in:"path" v:"required|min:1"`
}

type DetailRes struct {
	SlotItem
}

type UpdateReq struct {
	g.Meta   `path:"/admin/slots/{id}" method:"put" tags:"Admin/Slot" summary:"更新广告位"`
	Id       int64  `json:"id" in:"path" v:"required|min:1"`
	Name     string `json:"name" v:"required#name必填"`
	SlotType string `json:"slot_type" d:"banner"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Status   int    `json:"status" d:"1"`
	Remark   string `json:"remark"`
}

type UpdateRes struct {
	Id int64 `json:"id"`
}

type DeleteReq struct {
	g.Meta `path:"/admin/slots/{id}" method:"delete" tags:"Admin/Slot" summary:"停用广告位"`
	Id     int64 `json:"id" in:"path" v:"required|min:1"`
}

type DeleteRes struct {
	Id     int64 `json:"id"`
	Status int   `json:"status"`
}
