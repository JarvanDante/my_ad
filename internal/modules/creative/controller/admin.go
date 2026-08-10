package controller

import (
	"context"

	v1 "github.com/JarvanDante/my_ad/api/admin/creative/v1"
	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/modules/creative/logic"
)

type Admin struct {
	svc *logic.Service
}

func NewAdmin(svc *logic.Service) *Admin { return &Admin{svc: svc} }

func toItem(c dao.Creative) v1.CreativeItem {
	return v1.CreativeItem{
		Id: c.Id, Title: c.Title, MediaURL: c.MediaURL, LinkURL: c.LinkURL,
		StorageObjectId: c.StorageObjectId, Status: c.Status, Remark: c.Remark,
		CreatedAt: c.CreatedAt, UpdatedAt: c.UpdatedAt,
	}
}

func (c *Admin) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := c.svc.List(ctx, req.Page, req.Size, req.Keyword, req.Status)
	if err != nil {
		return nil, err
	}
	res = &v1.ListRes{Total: total, List: make([]v1.CreativeItem, 0, len(list))}
	for _, x := range list {
		res.List = append(res.List, toItem(x))
	}
	return res, nil
}

func (c *Admin) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := c.svc.Create(ctx, &dao.Creative{
		Title: req.Title, MediaURL: req.MediaURL, LinkURL: req.LinkURL,
		StorageObjectId: req.StorageObjectId, Status: req.Status, Remark: req.Remark,
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
	return &v1.DetailRes{CreativeItem: toItem(*item)}, nil
}

func (c *Admin) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	if err := c.svc.Update(ctx, &dao.Creative{
		Id: req.Id, Title: req.Title, MediaURL: req.MediaURL, LinkURL: req.LinkURL,
		StorageObjectId: req.StorageObjectId, Status: req.Status, Remark: req.Remark,
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
