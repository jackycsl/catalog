package data

import (
	"context"

	"github.com/jackycsl/catalog/catalog-job/internal/biz"
	"github.com/jackycsl/catalog/catalog-job/internal/data/ent"
	"github.com/jackycsl/catalog/pkg/util/helper"
	"github.com/jackycsl/catalog/pkg/util/pagination"
)

func (r *gameRepo) DbGetGame(ctx context.Context, id int64) (*biz.Game, error) {
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
		CreatedAt:   po.CreatedAt,
	}, nil
}

func (r *gameRepo) DbCreateGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
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

func (r *gameRepo) DbListGame(ctx context.Context, pageNum, pageSize int64) ([]*biz.Game, error) {
	pos, err := r.data.db.Game.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
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

func (r *gameRepo) DbUpdateGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
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
