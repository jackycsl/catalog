package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	ctV1 "github.com/jackycsl/catalog/api/catalog/service/v1"
	"github.com/jackycsl/catalog/bff/interface/internal/biz"
)

var _ biz.CatalogRepo = (*catalogRepo)(nil)

type catalogRepo struct {
	data *Data
	log  *log.Helper
}

func NewCatalogRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *catalogRepo) GetGame(ctx context.Context, id int64) (*biz.Game, error) {
	reply, err := r.data.cc.GetGame(ctx, &ctV1.GetGameReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Game{
		Id:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Count:       reply.Count,
	}, err
}

func (r *catalogRepo) ListGame(ctx context.Context, pageNum, pageSize int64) ([]*biz.Game, error) {
	reply, err := r.data.cc.ListGame(ctx, &ctV1.ListGameReq{
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Game, 0)
	for _, x := range reply.Results {
		rv = append(rv, &biz.Game{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
		})
	}
	return rv, nil
}
