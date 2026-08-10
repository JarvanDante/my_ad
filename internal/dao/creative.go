package dao

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Creative struct {
	Id              string
	Title           string
	MediaURL        string
	LinkURL         string
	StorageObjectId string
	Status          int
	Remark          string
	CreatedAt       string
	UpdatedAt       string
}

type CreativeRepo struct{}

func NewCreativeRepo() *CreativeRepo { return &CreativeRepo{} }

func (r *CreativeRepo) List(ctx context.Context, page, size int, keyword string, status int) ([]Creative, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	m := g.DB().Model("ad_creative").Ctx(ctx).Safe()
	if keyword != "" {
		kw := "%" + keyword + "%"
		m = m.Where("(id LIKE ? OR title LIKE ?)", kw, kw)
	}
	if status == 0 || status == 1 {
		m = m.Where("status", status)
	}
	total, err := m.Clone().Count()
	if err != nil {
		return nil, 0, err
	}
	rows, err := m.OrderDesc("created_at").Page(page, size).All()
	if err != nil {
		return nil, 0, err
	}
	list := make([]Creative, 0, len(rows))
	for _, row := range rows {
		list = append(list, mapCreative(row.Map()))
	}
	return list, total, nil
}

func (r *CreativeRepo) Create(ctx context.Context, c *Creative) error {
	_, err := g.DB().Model("ad_creative").Ctx(ctx).Data(g.Map{
		"id": c.Id, "title": c.Title, "media_url": c.MediaURL, "link_url": c.LinkURL,
		"storage_object_id": c.StorageObjectId, "status": c.Status, "remark": c.Remark,
		"created_at": gtime.Now(), "updated_at": gtime.Now(),
	}).Insert()
	return err
}

func (r *CreativeRepo) Update(ctx context.Context, c *Creative) error {
	_, err := g.DB().Model("ad_creative").Ctx(ctx).Where("id", c.Id).Data(g.Map{
		"title": c.Title, "media_url": c.MediaURL, "link_url": c.LinkURL,
		"storage_object_id": c.StorageObjectId, "status": c.Status, "remark": c.Remark,
		"updated_at": gtime.Now(),
	}).Update()
	return err
}

func (r *CreativeRepo) Get(ctx context.Context, id string) (*Creative, error) {
	row, err := g.DB().Model("ad_creative").Ctx(ctx).Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, nil
	}
	c := mapCreative(row.Map())
	return &c, nil
}

func (r *CreativeRepo) SoftDisable(ctx context.Context, id string) error {
	_, err := g.DB().Model("ad_creative").Ctx(ctx).Where("id", id).Data(g.Map{
		"status": 0, "updated_at": gtime.Now(),
	}).Update()
	return err
}

func mapCreative(m g.Map) Creative {
	return Creative{
		Id: g.NewVar(m["id"]).String(), Title: g.NewVar(m["title"]).String(),
		MediaURL: g.NewVar(m["media_url"]).String(), LinkURL: g.NewVar(m["link_url"]).String(),
		StorageObjectId: g.NewVar(m["storage_object_id"]).String(),
		Status:          g.NewVar(m["status"]).Int(), Remark: g.NewVar(m["remark"]).String(),
		CreatedAt: g.NewVar(m["created_at"]).String(), UpdatedAt: g.NewVar(m["updated_at"]).String(),
	}
}
