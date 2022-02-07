package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackycsl/catalog/catalog-job/internal/biz"
)

var _ biz.CatalogJobRepo = (*catalogJobRepo)(nil)

type catalogJobRepo struct {
	data *Data
	log  *log.Helper
}

func NewCatalogJobRepo(data *Data, logger log.Logger) biz.CatalogJobRepo {
	return &catalogJobRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "data/catalog")),
	}
}
