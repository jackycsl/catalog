package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type CatalogJobRepo interface {
}

type CatalogJobUseCase struct {
	repo CatalogJobRepo
	log  *log.Helper
}

func NewCatalogJobUseCase(repo CatalogJobRepo, logger log.Logger) *CatalogJobUseCase {
	return &CatalogJobUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/catalog/job"))}
}
