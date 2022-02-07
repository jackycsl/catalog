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

type CatalogJobRepo interface {
	// Redis
	CacheSetGame(ctx context.Context, c *Game) (*Game, error)
	// Kafka
	KafkaConsumeBackfillGame(ctx context.Context) (*Game, error)
}

type CatalogJobUseCase struct {
	repo CatalogJobRepo
	log  *log.Helper
}

func NewCatalogJobUseCase(repo CatalogJobRepo, logger log.Logger) *CatalogJobUseCase {
	return &CatalogJobUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/catalog/job"))}
}

func (uc *CatalogJobUseCase) Consume(ctx context.Context) (*Game, error) {
	return uc.repo.KafkaConsumeBackfillGame(ctx)
}
