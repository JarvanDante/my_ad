package v1

import "github.com/gogf/gf/v2/frame/g"

type AdItem struct {
	CampaignId string `json:"campaign_id"`
	CreativeId string `json:"creative_id"`
	Title      string `json:"title"`
	MediaURL   string `json:"media_url"`
	LinkURL    string `json:"link_url"`
	SlotCode   string `json:"slot_code"`
	Priority   int    `json:"priority"`
	Weight     int    `json:"weight"`
}

type FetchReq struct {
	g.Meta   `path:"/open/ads" method:"get" tags:"Open/Ad" summary:"按广告位拉取可展示广告"`
	SlotCode string `json:"slot_code" v:"required#slot_code必填"`
	Limit    int    `json:"limit" d:"10"`
}

type FetchRes struct {
	List []AdItem `json:"list"`
}

type EventReq struct {
	g.Meta     `path:"/open/events" method:"post" tags:"Open/Ad" summary:"曝光/点击上报"`
	EventType  string `json:"event_type" v:"required|in:impression,click"`
	CampaignId string `json:"campaign_id"`
	CreativeId string `json:"creative_id"`
	SlotCode   string `json:"slot_code"`
}

type EventRes struct {
	Ok bool `json:"ok"`
}
