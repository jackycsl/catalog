package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/jackycsl/catalog/api/bff/interface/v1"
	"github.com/jackycsl/catalog/bff/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	v1.UnimplementedShopInterfaceServer

	cc  *biz.CatalogUsecase
	log *log.Helper
}

func NewShopInterface(
	cc *biz.CatalogUsecase,
	logger log.Logger) *ShopInterface {
	return &ShopInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		cc:  cc,
	}
}
