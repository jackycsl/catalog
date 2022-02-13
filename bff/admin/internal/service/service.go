package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/jackycsl/catalog/api/bff/admin/v1"
	"github.com/jackycsl/catalog/bff/admin/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopAdmin)

type ShopAdmin struct {
	v1.UnimplementedShopAdminServer

	log *log.Helper
	cc  *biz.CatalogUsecase
}

func NewShopAdmin(cc *biz.CatalogUsecase, logger log.Logger) *ShopAdmin {
	return &ShopAdmin{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		cc:  cc,
	}
}
