package controller

import (
	"context"

	v1 "github.com/JarvanDante/my_ad/api/admin/campaign/v1"
	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/modules/campaign/logic"
)

type Admin struct {
	svc *logic.Service
}

func NewAdmin(svc *logic.Service) *Admin { return &Admin{svc: svc} }

func toItem(c dao.Campaign) v1.CampaignItem {
	return v1.CampaignItem{
		Id: c.Id, Name: c.Name, SlotId: c.SlotId, CreativeId: c.CreativeId,
		SiteCode: c.SiteCode, Priority: c.Priority, Weight: c.Weight, Status: c.Status,
		StartAt: c.StartAt, EndAt: c.EndAt, Remark: c.Remark,
		CreatedAt: c.CreatedAt, UpdatedAt: c.UpdatedAt,
	}
}

func (c *Admin) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := c.svc.List(ctx, req.Page, req.Size, req.Keyword, req.SiteCode, req.SlotId, req.Status)
	if err != nil {
		return nil, err
	}
	res = &v1.ListRes{Total: total, List: make([]v1.CampaignItem, 0, len(list))}
	for _, x := range list {
		res.List = append(res.List, toItem(x))
	}
	return res, nil
}

func (c *Admin) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := c.svc.Create(ctx, &dao.Campaign{
		Name: req.Name, SlotId: req.SlotId, CreativeId: req.CreativeId, SiteCode: req.SiteCode,
		Priority: req.Priority, Weight: req.Weight, Status: req.Status,
		StartAt: req.StartAt, EndAt: req.EndAt, Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id}, nil
}

func (c *Admin) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	item, err := c.svc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DetailRes{CampaignItem: toItem(*item)}, nil
}

func (c *Admin) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	if err := c.svc.Update(ctx, &dao.Campaign{
		Id: req.Id, Name: req.Name, SlotId: req.SlotId, CreativeId: req.CreativeId,
		SiteCode: req.SiteCode, Priority: req.Priority, Weight: req.Weight, Status: req.Status,
		StartAt: req.StartAt, EndAt: req.EndAt, Remark: req.Remark,
	}); err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Id: req.Id}, nil
}

func (c *Admin) SetStatus(ctx context.Context, req *v1.SetStatusReq) (res *v1.SetStatusRes, err error) {
	if err := c.svc.SetStatus(ctx, req.Id, req.Status); err != nil {
		return nil, err
	}
	return &v1.SetStatusRes{Id: req.Id, Status: req.Status}, nil
}
