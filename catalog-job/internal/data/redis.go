package data

import (
	"context"
	"fmt"
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

	// r.data.rdb.ZAdd(ctx, "game_index_cache:", &redis.Z{Score: float64(b.CreatedAt.Unix()), Member: game})
	// r.data.rdb.Expire(ctx, "game_index_cache:", 8*time.Hour)

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

func (r *gameRepo) CacheSetGameList(ctx context.Context, gs []*biz.Game) error {

	pipe := r.data.rdb.Pipeline()

	for _, g := range gs {
		gameId := strconv.Itoa(int(g.Id))
		game := "game_content:" + gameId

		pipe.HSetNX(ctx, game, "Name", g.Name)
		pipe.HSetNX(ctx, game, "Description", g.Description)
		pipe.HSetNX(ctx, game, "Count", g.Count)
		pipe.Expire(ctx, game, 8*time.Hour)

		pipe.ZAdd(ctx, "game_index_cache:", &redis.Z{Score: float64(g.CreatedAt.Unix()), Member: game})
		pipe.Expire(ctx, "game_index_cache:", 8*time.Hour)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
