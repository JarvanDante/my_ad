package logic

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/JarvanDante/my_ad/internal/dao"
	"github.com/JarvanDante/my_ad/internal/shared/errcode"
	"github.com/JarvanDante/my_ad/internal/shared/kit"
)

type Service struct {
	repo *dao.CreativeRepo
}

func New(repo *dao.CreativeRepo) *Service { return &Service{repo: repo} }

func (s *Service) List(ctx context.Context, page, size int, keyword string, status int) ([]dao.Creative, int, error) {
	return s.repo.List(ctx, page, size, keyword, status)
}

func (s *Service) Create(ctx context.Context, in *dao.Creative) (string, error) {
	id, err := kit.NewPublicID()
	if err != nil {
		return "", err
	}
	in.Id = id
	if in.Status != 0 && in.Status != 1 {
		in.Status = 1
	}
	if err := s.repo.Create(ctx, in); err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) Get(ctx context.Context, id string) (*dao.Creative, error) {
	c, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, gerror.NewCode(errcode.CodeNotFound, "素材不存在")
	}
	return c, nil
}

func (s *Service) Update(ctx context.Context, in *dao.Creative) error {
	if _, err := s.Get(ctx, in.Id); err != nil {
		return err
	}
	return s.repo.Update(ctx, in)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if _, err := s.Get(ctx, id); err != nil {
		return err
	}
	return s.repo.SoftDisable(ctx, id)
}
