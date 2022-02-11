package data

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackycsl/catalog/catalog-service/internal/biz"
	"github.com/jackycsl/catalog/pkg/util/helper"
	"github.com/jackycsl/catalog/pkg/util/pagination"
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
	order := "game_index_cache:"
	start := pagination.GetPageOffset(pageNum, pageSize)
	end := start + pageSize - 1

	ids := r.data.rdb.ZRange(ctx, order, start, end).Val()

	m := map[string]*redis.StringStringMapCmd{}
	pipe := r.data.rdb.Pipeline()
	for _, id := range ids {
		m[id] = pipe.HGetAll(ctx, id)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Println(err)
	}

	if len(m) == 0 {
		return nil, helper.ErrRecordNotFound
	}

	rv := make([]*biz.Game, 0)
	for k, v := range m {
		gameData, _ := v.Result()
		id, _ := strconv.ParseInt(strings.Split(k, ":")[1], 10, 64)
		count, _ := strconv.ParseInt(gameData["Count"], 10, 64)
		rv = append(rv, &biz.Game{
			Id:          id,
			Name:        gameData["Name"],
			Description: gameData["Description"],
			Count:       count,
		})
	}

	return rv, nil
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
