package logic

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/shared/errcode"
	"github.com/JarvanDante/my_ad/internal/shared/kit"
)

type Service struct {
	repo      *dao.CampaignRepo
	slots     *dao.SlotRepo
	creatives *dao.CreativeRepo
}

func New(repo *dao.CampaignRepo, slots *dao.SlotRepo, creatives *dao.CreativeRepo) *Service {
	return &Service{repo: repo, slots: slots, creatives: creatives}
}

func (s *Service) List(ctx context.Context, page, size int, keyword, siteCode string, slotId int64, status int) ([]dao.Campaign, int, error) {
	return s.repo.List(ctx, page, size, keyword, siteCode, slotId, status)
}

func (s *Service) Create(ctx context.Context, in *dao.Campaign) (string, error) {
	if err := s.validateRefs(ctx, in.SlotId, in.CreativeId); err != nil {
		return "", err
	}
	id, err := kit.NewPublicID()
	if err != nil {
		return "", err
	}
	in.Id = id
	in.SiteCode = strings.ToUpper(strings.TrimSpace(in.SiteCode))
	if in.Status != 0 && in.Status != 1 {
		in.Status = 1
	}
	if in.Priority <= 0 {
		in.Priority = 100
	}
	if in.Weight <= 0 {
		in.Weight = 100
	}
	if err := s.repo.Create(ctx, in); err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) Get(ctx context.Context, id string) (*dao.Campaign, error) {
	c, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, gerror.NewCode(errcode.CodeNotFound, "投放不存在")
	}
	return c, nil
}

func (s *Service) Update(ctx context.Context, in *dao.Campaign) error {
	if _, err := s.Get(ctx, in.Id); err != nil {
		return err
	}
	if err := s.validateRefs(ctx, in.SlotId, in.CreativeId); err != nil {
		return err
	}
	in.SiteCode = strings.ToUpper(strings.TrimSpace(in.SiteCode))
	return s.repo.Update(ctx, in)
}

func (s *Service) SetStatus(ctx context.Context, id string, status int) error {
	if _, err := s.Get(ctx, id); err != nil {
		return err
	}
	return s.repo.SetStatus(ctx, id, status)
}

func (s *Service) Fetch(ctx context.Context, slotCode, siteCode string, limit int) ([]dao.FetchItem, error) {
	return s.repo.FetchForSite(ctx, slotCode, strings.ToUpper(strings.TrimSpace(siteCode)), limit)
}

func (s *Service) ReportEvent(ctx context.Context, eventType, campaignId, creativeId, slotCode, siteCode, appKey string) error {
	return s.repo.InsertEvent(ctx, eventType, campaignId, creativeId, slotCode, siteCode, appKey)
}

func (s *Service) validateRefs(ctx context.Context, slotId int64, creativeId string) error {
	slot, err := s.slots.Get(ctx, slotId)
	if err != nil {
		return err
	}
	if slot == nil {
		return gerror.NewCode(errcode.CodeBadRequest, "广告位不存在")
	}
	cr, err := s.creatives.Get(ctx, creativeId)
	if err != nil {
		return err
	}
	if cr == nil {
		return gerror.NewCode(errcode.CodeBadRequest, "素材不存在")
	}
	return nil
}
