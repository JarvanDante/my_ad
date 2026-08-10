package dao

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Slot struct {
	Id        int64
	Code      string
	Name      string
	SlotType  string
	Width     int
	Height    int
	Status    int
	Remark    string
	CreatedAt string
	UpdatedAt string
}

type SlotRepo struct{}

func NewSlotRepo() *SlotRepo { return &SlotRepo{} }

func (r *SlotRepo) List(ctx context.Context, page, size int, keyword string, status int) ([]Slot, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	m := g.DB().Model("ad_slot").Ctx(ctx).Safe()
	if keyword != "" {
		kw := "%" + keyword + "%"
		m = m.Where("(code LIKE ? OR name LIKE ?)", kw, kw)
	}
	if status == 0 || status == 1 {
		m = m.Where("status", status)
	}
	total, err := m.Clone().Count()
	if err != nil {
		return nil, 0, err
	}
	rows, err := m.OrderDesc("id").Page(page, size).All()
	if err != nil {
		return nil, 0, err
	}
	list := make([]Slot, 0, len(rows))
	for _, row := range rows {
		list = append(list, mapSlot(row.Map()))
	}
	return list, total, nil
}

func (r *SlotRepo) Create(ctx context.Context, s *Slot) (int64, error) {
	res, err := g.DB().Model("ad_slot").Ctx(ctx).Data(g.Map{
		"code": s.Code, "name": s.Name, "slot_type": s.SlotType,
		"width": s.Width, "height": s.Height, "status": s.Status,
		"remark": s.Remark, "created_at": gtime.Now(), "updated_at": gtime.Now(),
	}).Insert()
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return id, nil
}

func (r *SlotRepo) Update(ctx context.Context, s *Slot) error {
	_, err := g.DB().Model("ad_slot").Ctx(ctx).Where("id", s.Id).Data(g.Map{
		"name": s.Name, "slot_type": s.SlotType, "width": s.Width, "height": s.Height,
		"status": s.Status, "remark": s.Remark, "updated_at": gtime.Now(),
	}).Update()
	return err
}

func (r *SlotRepo) Get(ctx context.Context, id int64) (*Slot, error) {
	row, err := g.DB().Model("ad_slot").Ctx(ctx).Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, nil
	}
	s := mapSlot(row.Map())
	return &s, nil
}

func (r *SlotRepo) GetByCode(ctx context.Context, code string) (*Slot, error) {
	row, err := g.DB().Model("ad_slot").Ctx(ctx).Where("code", strings.TrimSpace(code)).One()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return nil, nil
	}
	s := mapSlot(row.Map())
	return &s, nil
}

func (r *SlotRepo) SoftDisable(ctx context.Context, id int64) error {
	_, err := g.DB().Model("ad_slot").Ctx(ctx).Where("id", id).Data(g.Map{
		"status": 0, "updated_at": gtime.Now(),
	}).Update()
	return err
}

func mapSlot(m g.Map) Slot {
	return Slot{
		Id: g.NewVar(m["id"]).Int64(), Code: g.NewVar(m["code"]).String(),
		Name: g.NewVar(m["name"]).String(), SlotType: g.NewVar(m["slot_type"]).String(),
		Width: g.NewVar(m["width"]).Int(), Height: g.NewVar(m["height"]).Int(),
		Status: g.NewVar(m["status"]).Int(), Remark: g.NewVar(m["remark"]).String(),
		CreatedAt: g.NewVar(m["created_at"]).String(), UpdatedAt: g.NewVar(m["updated_at"]).String(),
	}
}
