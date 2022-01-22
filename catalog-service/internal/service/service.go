package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/jackycsl/catalog/api/catalog/service/v1"
	"github.com/jackycsl/catalog/catalog-service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogService)

type CatalogService struct {
	v1.UnimplementedCatalogServer

	bc  *biz.GameUseCase
	log *log.Helper
}

func NewCatalogService(bc *biz.GameUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}
