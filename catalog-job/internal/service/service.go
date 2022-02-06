package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/jackycsl/catalog/api/catalog/job/v1"

	"github.com/jackycsl/catalog/catalog-job/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogJob)

type CatalogJob struct {
	v1.UnimplementedCatalogJobServer

	oc  *biz.CatalogJobUseCase
	log *log.Helper
}

func NewCatalogJob(oc *biz.CatalogJobUseCase, logger log.Logger) *CatalogJob {
	return &CatalogJob{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "job/catalog"))}
}
