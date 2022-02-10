package data

import (
	"context"

	"github.com/jackycsl/catalog/catalog-job/internal/biz"
	"github.com/jackycsl/catalog/catalog-job/internal/data/ent"
	"github.com/jackycsl/catalog/pkg/util/helper"
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
