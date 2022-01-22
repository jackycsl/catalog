package data

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackycsl/catalog/catalog-service/internal/biz"
)

func idKey(id int64) string {
	return fmt.Sprintf("%d", id)
}

func (r *gameRepo) CacheGetGame(ctx context.Context, id int64) (*biz.Game, error) {

	gameData := r.data.rdb.HGetAll(ctx, idKey(id)).Val()

	i, _ := strconv.ParseInt(gameData["Id"], 10, 64)
	count, _ := strconv.ParseInt(gameData["Count"], 10, 64)

	return &biz.Game{
		Id:          i,
		Name:        gameData["Name"],
		Description: gameData["Description"],
		Count:       count,
		Images:      append([]biz.Image{}, biz.Image{gameData["images"]}),
	}, nil
}

func (r *gameRepo) CacheListGame(ctx context.Context, pageNum int64, pageSize int64) ([]*biz.Game, error) {
	return nil, nil
}
