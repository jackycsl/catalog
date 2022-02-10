package data

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackycsl/catalog/catalog-job/internal/biz"
)

func (r *gameRepo) CacheSetGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
	gameId := strconv.Itoa(int(b.Id))

	game := "game_content:" + gameId

	r.data.rdb.HSetNX(ctx, game, "Name", b.Name)
	r.data.rdb.HSetNX(ctx, game, "Description", b.Description)
	r.data.rdb.HSetNX(ctx, game, "Count", b.Count)
	r.data.rdb.Expire(ctx, game, 8*time.Hour)

	r.data.rdb.ZAdd(ctx, "game_index_cache:", &redis.Z{Score: float64(b.CreatedAt.Unix()), Member: game})
	r.data.rdb.Expire(ctx, "game_index_cache:", 8*time.Hour)

	gameData := r.data.rdb.HGetAll(ctx, game).Val()
	if len(gameData) == 0 {
		return nil, nil
	}

	i, _ := strconv.ParseInt(gameId, 10, 64)
	count, _ := strconv.ParseInt(gameData["Count"], 10, 64)

	return &biz.Game{
		Id:          i,
		Name:        gameData["Name"],
		Description: gameData["Description"],
		Count:       count,
	}, nil
}

func (r *gameRepo) CacheCreateGame(ctx context.Context, b *biz.Game) (*biz.Game, error) {
	gameId := strconv.Itoa(int(b.Id))

	game := "game_content:" + gameId
	r.data.rdb.HSet(ctx, game, map[string]interface{}{
		"Name":        b.Name,
		"Description": b.Description,
		"Count":       b.Count,
		"Created_at":  b.CreatedAt,
	})
	r.data.rdb.Expire(ctx, game, 8*time.Hour)

	r.data.rdb.ZAdd(ctx, "game_index_cache:", &redis.Z{Score: float64(b.CreatedAt.Unix()), Member: game})
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
