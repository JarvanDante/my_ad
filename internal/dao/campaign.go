package dao

import (
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Campaign struct {
	Id         string
	Name       string
	SlotId     int64
	CreativeId string
	SiteCode   string
	Priority   int
	Weight     int
	Status     int
	StartAt    string
	EndAt      string
	Remark     string
	CreatedAt  string
	UpdatedAt  string
}

type FetchItem struct {
	CampaignId string
	CreativeId string
	Title      string
	MediaURL   string
	LinkURL    string
	SlotCode   string
	Priority   int
	Weight     int
}

type CampaignRepo struct{}

func NewCampaignRepo() *CampaignRepo { return &CampaignRepo{} }

func (r *CampaignRepo) List(ctx context.Context, page, size int, keyword, siteCode string, slotId int64, status int) ([]Campaign, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	m := g.DB().Model("ad_campaign").Ctx(ctx).Safe()
	if keyword != "" {
		kw := "%" + keyword + "%"
		m = m.Where("(id LIKE ? OR name LIKE ?)", kw, kw)
	}
	if siteCode != "" {
		m = m.Where("site_code", siteCode)
	}
	if slotId > 0 {
		m = m.Where("slot_id", slotId)
	}
	if status == 0 || status == 1 {
		m = m.Where("status", status)
	}
	total, err := m.Clone().Count()
	if err != nil {
		return nil, 0, err
	}
	rows, err := m.OrderDesc("priority").OrderDesc("created_at").Page(page, size).All()
	if err != nil {
		return nil, 0, err
	}
	list := make([]Campaign, 0, len(rows))
	for _, row := range rows {
		list = append(list, mapCampaign(row.Map()))
	}
	return list, total, nil
}

func (r *CampaignRepo) Create(ctx context.Context, c *Campaign) error {
	_, err := g.DB().Model("ad_campaign").Ctx(ctx).Data(campaignData(c, true)).Insert()
	return err
}

func (r *CampaignRepo) Update(ctx context.Context, c *Campaign) error {
	_, err := g.DB().Model("ad_campaign").Ctx(ctx).Where("id", c.Id).Data(campaignData(c, false)).Update()
	return err
}

func (r *CampaignRepo) SetStatus(ctx context.Context, id string, status int) error {
	_, err := g.DB().Model("ad_campaign").Ctx(ctx).Where("id", id).Data(g.Map{
		"status": status, "updated_at": gtime.Now(),
	}).Update()
	return err
}

func (r *CampaignRepo) Get(ctx context.Context, id string) (*Campaign, error) {
	row, err := g.DB().Model("ad_campaign").Ctx(ctx).Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, nil
	}
	c := mapCampaign(row.Map())
	return &c, nil
}

// FetchForSite 按广告位 + 站点拉取有效投放（全站投放 site_code 空 或精确匹配）。
func (r *CampaignRepo) FetchForSite(ctx context.Context, slotCode, siteCode string, limit int) ([]FetchItem, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}
	now := time.Now()
	sql := `
SELECT c.id AS campaign_id, c.creative_id, cr.title, cr.media_url, cr.link_url,
       s.code AS slot_code, c.priority, c.weight
FROM ad_campaign c
JOIN ad_slot s ON s.id = c.slot_id
JOIN ad_creative cr ON cr.id = c.creative_id
WHERE s.code = ?
  AND s.status = 1
  AND c.status = 1
  AND cr.status = 1
  AND (c.site_code = '' OR c.site_code = ?)
  AND (c.start_at IS NULL OR c.start_at <= ?)
  AND (c.end_at IS NULL OR c.end_at >= ?)
ORDER BY c.priority DESC, c.weight DESC, c.created_at DESC
LIMIT ?`
	rows, err := g.DB().Ctx(ctx).Raw(sql, strings.TrimSpace(slotCode), strings.TrimSpace(siteCode), now, now, limit).All()
	if err != nil {
		return nil, err
	}
	out := make([]FetchItem, 0, len(rows))
	for _, row := range rows {
		m := row.Map()
		out = append(out, FetchItem{
			CampaignId: g.NewVar(m["campaign_id"]).String(),
			CreativeId: g.NewVar(m["creative_id"]).String(),
			Title:      g.NewVar(m["title"]).String(),
			MediaURL:   g.NewVar(m["media_url"]).String(),
			LinkURL:    g.NewVar(m["link_url"]).String(),
			SlotCode:   g.NewVar(m["slot_code"]).String(),
			Priority:   g.NewVar(m["priority"]).Int(),
			Weight:     g.NewVar(m["weight"]).Int(),
		})
	}
	return out, nil
}

func (r *CampaignRepo) InsertEvent(ctx context.Context, eventType, campaignId, creativeId, slotCode, siteCode, appKey string) error {
	_, err := g.DB().Model("ad_event").Ctx(ctx).Data(g.Map{
		"event_type": eventType, "campaign_id": campaignId, "creative_id": creativeId,
		"slot_code": slotCode, "site_code": siteCode, "app_key": appKey,
		"created_at": gtime.Now(),
	}).Insert()
	return err
}

func campaignData(c *Campaign, create bool) g.Map {
	data := g.Map{
		"name": c.Name, "slot_id": c.SlotId, "creative_id": c.CreativeId,
		"site_code": c.SiteCode, "priority": c.Priority, "weight": c.Weight,
		"status": c.Status, "remark": c.Remark, "updated_at": gtime.Now(),
	}
	if c.StartAt != "" {
		data["start_at"] = c.StartAt
	} else {
		data["start_at"] = nil
	}
	if c.EndAt != "" {
		data["end_at"] = c.EndAt
	} else {
		data["end_at"] = nil
	}
	if create {
		data["id"] = c.Id
		data["created_at"] = gtime.Now()
	}
	return data
}

func mapCampaign(m g.Map) Campaign {
	return Campaign{
		Id: g.NewVar(m["id"]).String(), Name: g.NewVar(m["name"]).String(),
		SlotId: g.NewVar(m["slot_id"]).Int64(), CreativeId: g.NewVar(m["creative_id"]).String(),
		SiteCode: g.NewVar(m["site_code"]).String(), Priority: g.NewVar(m["priority"]).Int(),
		Weight: g.NewVar(m["weight"]).Int(), Status: g.NewVar(m["status"]).Int(),
		StartAt: g.NewVar(m["start_at"]).String(), EndAt: g.NewVar(m["end_at"]).String(),
		Remark:    g.NewVar(m["remark"]).String(),
		CreatedAt: g.NewVar(m["created_at"]).String(), UpdatedAt: g.NewVar(m["updated_at"]).String(),
	}
}
