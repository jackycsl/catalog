package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Game struct {
	Id          int64
	Name        string
	Description string
	Count       int64
}

type CatalogRepo interface {
	GetGame(ctx context.Context, id int64) (*Game, error)
	ListGame(ctx context.Context, pageNum, pageSize int64) ([]*Game, error)
}

type CatalogUsecase struct {
	repo CatalogRepo
	log  *log.Helper
}

func NewCatalogUsecase(repo CatalogRepo, logger log.Logger) *CatalogUsecase {
	return &CatalogUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CatalogUsecase) GetGame(ctx context.Context, id int64) (*Game, error) {
	return uc.repo.GetGame(ctx, id)
}

func (uc *CatalogUsecase) ListGame(ctx context.Context, pageNum, pageSize int64) ([]*Game, error) {
	return uc.repo.ListGame(ctx, pageNum, pageSize)
}
