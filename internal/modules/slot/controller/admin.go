package controller

import (
	"context"

	v1 "github.com/JarvanDante/my_ad/api/admin/slot/v1"
	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/modules/slot/logic"
)

type Admin struct {
	svc *logic.Service
}

func NewAdmin(svc *logic.Service) *Admin { return &Admin{svc: svc} }

func toItem(s dao.Slot) v1.SlotItem {
	return v1.SlotItem{
		Id: s.Id, Code: s.Code, Name: s.Name, SlotType: s.SlotType,
		Width: s.Width, Height: s.Height, Status: s.Status, Remark: s.Remark,
		CreatedAt: s.CreatedAt, UpdatedAt: s.UpdatedAt,
	}
}

func (c *Admin) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := c.svc.List(ctx, req.Page, req.Size, req.Keyword, req.Status)
	if err != nil {
		return nil, err
	}
	res = &v1.ListRes{Total: total, List: make([]v1.SlotItem, 0, len(list))}
	for _, x := range list {
		res.List = append(res.List, toItem(x))
	}
	return res, nil
}

func (c *Admin) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := c.svc.Create(ctx, &dao.Slot{
		Code: req.Code, Name: req.Name, SlotType: req.SlotType,
		Width: req.Width, Height: req.Height, Status: req.Status, Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id, Code: req.Code}, nil
}

func (c *Admin) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	slot, err := c.svc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DetailRes{SlotItem: toItem(*slot)}, nil
}

func (c *Admin) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	if err := c.svc.Update(ctx, &dao.Slot{
		Id: req.Id, Name: req.Name, SlotType: req.SlotType,
		Width: req.Width, Height: req.Height, Status: req.Status, Remark: req.Remark,
	}); err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Id: req.Id}, nil
}

func (c *Admin) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if err := c.svc.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &v1.DeleteRes{Id: req.Id, Status: 0}, nil
}
