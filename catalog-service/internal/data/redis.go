package data

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackycsl/catalog/catalog-service/internal/biz"
	"github.com/jackycsl/catalog/pkg/util/helper"
)

func idKey(id int64) string {
	return fmt.Sprintf("%d", id)
}

func (r *gameRepo) CacheGetGame(ctx context.Context, id int64) (*biz.Game, error) {

	game := "game_content:" + idKey(id)
	gameData := r.data.rdb.HGetAll(ctx, game).Val()

	if len(gameData) == 0 {
		return nil, helper.ErrRecordNotFound
	}

	count, _ := strconv.ParseInt(gameData["Count"], 10, 64)

	return &biz.Game{
		Id:          id,
		Name:        gameData["Name"],
		Description: gameData["Description"],
		Count:       count,
	}, nil
}

func (r *gameRepo) CacheListGame(ctx context.Context, pageNum int64, pageSize int64) ([]*biz.Game, error) {
	return nil, nil
}

func (r *gameRepo) CacheCreateGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
	gameId := strconv.Itoa(int(r.data.rdb.Incr(ctx, "game_id").Val()))

	now := time.Now().Unix()
	game := "game_content:" + gameId
	r.data.rdb.HSet(ctx, game, map[string]interface{}{
		"Name":        b.Name,
		"Description": b.Description,
		"Count":       b.Count,
		"Created_at":  now,
	})
	r.data.rdb.Expire(ctx, game, 8*time.Hour)

	r.data.rdb.ZAdd(ctx, "game_index_cache:", &redis.Z{Score: float64(now), Member: game})
	r.data.rdb.Expire(ctx, "game_index_cache:", 8*time.Hour)

	gameData := r.data.rdb.HGetAll(ctx, game).Val()

	i, _ := strconv.ParseInt(gameId, 10, 64)
	count, _ := strconv.ParseInt(gameData["Count"], 10, 64)

	return &biz.Game{
		Id:          i,
		Name:        gameData["Name"],
		Description: gameData["Description"],
		Count:       count,
	}, nil
}

func (r *gameRepo) CacheUpdateGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
	return nil, nil
}
