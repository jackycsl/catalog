package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackycsl/catalog/bff/admin/internal/biz"
)

var _ biz.CatalogRepo = (*catalogRepo)(nil)

type catalogRepo struct {
	data *Data
	log  *log.Helper
}

func NewCatalogRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/beer")),
	}
}

func (r *catalogRepo) GetGame(ctx context.Context, id int64) (*biz.Game, error) {
	return nil, nil
}

func (r *catalogRepo) ListGame(ctx context.Context, pageNum, pageSize int64) ([]*biz.Game, error) {
	return nil, nil
}
