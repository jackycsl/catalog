package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	ctV1 "github.com/jackycsl/catalog/api/catalog/service/v1"
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
		log:  log.NewHelper(log.With(logger, "module", "data/game")),
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
			Count:       x.Count,
		})
	}
	return rv, nil
}

func (r *catalogRepo) CreateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	reply, err := r.data.cc.CreateGame(ctx, &ctV1.CreateGameReq{
		Name:        c.Name,
		Description: c.Description,
		Count:       c.Count,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Game{
		Id:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Count:       reply.Count,
	}, nil
}

func (r *catalogRepo) UpdateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	reply, err := r.data.cc.UpdateGame(ctx, &ctV1.UpdateGameReq{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
		Count:       c.Count,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Game{
		Id:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Count:       reply.Count,
	}, nil
}
