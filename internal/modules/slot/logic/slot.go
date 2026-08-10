package logic

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/shared/errcode"
)

type Service struct {
	repo *dao.SlotRepo
}

func New(repo *dao.SlotRepo) *Service { return &Service{repo: repo} }

func (s *Service) List(ctx context.Context, page, size int, keyword string, status int) ([]dao.Slot, int, error) {
	return s.repo.List(ctx, page, size, keyword, status)
}

func (s *Service) Create(ctx context.Context, in *dao.Slot) (int64, error) {
	in.Code = strings.TrimSpace(in.Code)
	if in.Code == "" {
		return 0, gerror.NewCode(errcode.CodeBadRequest, "code 不能为空")
	}
	exist, err := s.repo.GetByCode(ctx, in.Code)
	if err != nil {
		return 0, err
	}
	if exist != nil {
		return 0, gerror.NewCode(errcode.CodeBadRequest, "广告位 code 已存在")
	}
	if in.SlotType == "" {
		in.SlotType = "banner"
	}
	if in.Status != 0 && in.Status != 1 {
		in.Status = 1
	}
	return s.repo.Create(ctx, in)
}

func (s *Service) Get(ctx context.Context, id int64) (*dao.Slot, error) {
	slot, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if slot == nil {
		return nil, gerror.NewCode(errcode.CodeNotFound, "广告位不存在")
	}
	return slot, nil
}

func (s *Service) Update(ctx context.Context, in *dao.Slot) error {
	if _, err := s.Get(ctx, in.Id); err != nil {
		return err
	}
	if in.SlotType == "" {
		in.SlotType = "banner"
	}
	return s.repo.Update(ctx, in)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if _, err := s.Get(ctx, id); err != nil {
		return err
	}
	return s.repo.SoftDisable(ctx, id)
}
