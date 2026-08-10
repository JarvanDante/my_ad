package v1

import "github.com/gogf/gf/v2/frame/g"

type CreativeItem struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	MediaURL        string `json:"media_url"`
	LinkURL         string `json:"link_url"`
	StorageObjectId string `json:"storage_object_id"`
	Status          int    `json:"status"`
	Remark          string `json:"remark"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type ListReq struct {
	g.Meta  `path:"/admin/creatives" method:"get" tags:"Admin/Creative" summary:"素材列表"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	Keyword string `json:"keyword"`
	Status  int    `json:"status" d:"-1"`
}

type ListRes struct {
	List  []CreativeItem `json:"list"`
	Total int            `json:"total"`
}

type CreateReq struct {
	g.Meta          `path:"/admin/creatives" method:"post" tags:"Admin/Creative" summary:"创建素材"`
	Title           string `json:"title" v:"required#title必填"`
	MediaURL        string `json:"media_url" v:"required#media_url必填"`
	LinkURL         string `json:"link_url"`
	StorageObjectId string `json:"storage_object_id"`
	Status          int    `json:"status" d:"1"`
	Remark          string `json:"remark"`
}

type CreateRes struct {
	Id string `json:"id"`
}

type DetailReq struct {
	g.Meta `path:"/admin/creatives/{id}" method:"get" tags:"Admin/Creative" summary:"素材详情"`
	Id     string `json:"id" in:"path" v:"required"`
}

type DetailRes struct {
	CreativeItem
}

type UpdateReq struct {
	g.Meta          `path:"/admin/creatives/{id}" method:"put" tags:"Admin/Creative" summary:"更新素材"`
	Id              string `json:"id" in:"path" v:"required"`
	Title           string `json:"title" v:"required#title必填"`
	MediaURL        string `json:"media_url" v:"required#media_url必填"`
	LinkURL         string `json:"link_url"`
	StorageObjectId string `json:"storage_object_id"`
	Status          int    `json:"status" d:"1"`
	Remark          string `json:"remark"`
}

type UpdateRes struct {
	Id string `json:"id"`
}

type DeleteReq struct {
	g.Meta `path:"/admin/creatives/{id}" method:"delete" tags:"Admin/Creative" summary:"下架素材"`
	Id     string `json:"id" in:"path" v:"required"`
}

type DeleteRes struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}
