package data

import (
	"context"

	"github.com/jackycsl/catalog/catalog-service/internal/biz"
	"github.com/jackycsl/catalog/catalog-service/internal/data/ent"
	"github.com/jackycsl/catalog/pkg/util/helper"
	"github.com/jackycsl/catalog/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.GameRepo = (*gameRepo)(nil)

type gameRepo struct {
	data *Data
	log  *log.Helper
}

func NewGameRepo(data *Data, logger log.Logger) biz.GameRepo {
	return &gameRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/game")),
	}
}

func (r *gameRepo) CreateGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
	po, err := r.data.db.Game.
		Create().
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Game{
		Id:          po.ID,
		Name:        po.Name,
		Description: po.Description,
		Count:       po.Count,
	}, nil
}

func (r *gameRepo) GetGame(ctx context.Context, id int64) (*biz.Game, error) {
	po, err := r.data.db.Game.Get(ctx, id)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, helper.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &biz.Game{
		Id:          po.ID,
		Name:        po.Name,
		Description: po.Description,
		Count:       po.Count,
	}, nil
}

func (r *gameRepo) UpdateGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
	po, err := r.data.db.Game.
		UpdateOneID(b.Id).
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, helper.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &biz.Game{
		Id:          po.ID,
		Name:        po.Name,
		Description: po.Description,
		Count:       po.Count,
	}, nil
}

func (r *gameRepo) ListGame(ctx context.Context, pageNum, pageSize int64) ([]*biz.Game, error) {
	pos, err := r.data.db.Game.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pos) == 0 {
		return nil, helper.ErrRecordNotFound
	}
	rv := make([]*biz.Game, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Game{
			Id:          po.ID,
			Name:        po.Name,
			Description: po.Description,
			Count:       po.Count,
		})
	}
	return rv, nil
}
