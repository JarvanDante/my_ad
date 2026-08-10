package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/JarvanDante/my_ad/api/open/ad/v1"
	"github.com/JarvanDante/my_ad/internal/modules/campaign/logic"
)

type Open struct {
	svc *logic.Service
}

func New(svc *logic.Service) *Open { return &Open{svc: svc} }

func (c *Open) Fetch(ctx context.Context, req *v1.FetchReq) (res *v1.FetchRes, err error) {
	siteCode := g.RequestFromCtx(ctx).GetCtxVar("site_code").String()
	list, err := c.svc.Fetch(ctx, req.SlotCode, siteCode, req.Limit)
	if err != nil {
		return nil, err
	}
	res = &v1.FetchRes{List: make([]v1.AdItem, 0, len(list))}
	for _, x := range list {
		res.List = append(res.List, v1.AdItem{
			CampaignId: x.CampaignId, CreativeId: x.CreativeId, Title: x.Title,
			MediaURL: x.MediaURL, LinkURL: x.LinkURL, SlotCode: x.SlotCode,
			Priority: x.Priority, Weight: x.Weight,
		})
	}
	return res, nil
}

func (c *Open) Event(ctx context.Context, req *v1.EventReq) (res *v1.EventRes, err error) {
	r := g.RequestFromCtx(ctx)
	siteCode := r.GetCtxVar("site_code").String()
	appKey := r.GetCtxVar("app_key").String()
	if err := c.svc.ReportEvent(ctx, req.EventType, req.CampaignId, req.CreativeId, req.SlotCode, siteCode, appKey); err != nil {
		return nil, err
	}
	return &v1.EventRes{Ok: true}, nil
}
