package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackycsl/catalog/catalog-job/internal/data/ent"
	"github.com/jackycsl/catalog/pkg/event/event"
	"github.com/jackycsl/catalog/pkg/util/helper"
)

type Game struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	CreatedAt   time.Time
}

type GameEntry struct {
	GameId int64 `json:"game_id"`
}

func Receive(receiver event.Receiver, db *ent.Client, rdb *redis.Client) error {
	var ctx = context.Background()
	fmt.Println("start receiver")

	err := receiver.Receive(ctx, func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var game GameEntry
		if err := json.Unmarshal(msg.Value(), &game); err != nil {
			log.Println(err)
		}
		po, err := db.Game.Get(ctx, game.GameId)
		if err != nil {
			switch {
			case ent.IsNotFound(err):
				return helper.ErrRecordNotFound
			default:
				return err
			}
		}
		ge := &Game{
			Id:          po.ID,
			Name:        po.Name,
			Description: po.Description,
			Count:       po.Count,
			CreatedAt:   po.CreatedAt,
		}
		g, err := cacheSetGame(ctx, ge, rdb)
		if err != nil {
			return nil
		}
		fmt.Printf("Cache Set, key: %d\n", g.Id)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func cacheSetGame(ctx context.Context, b *Game, rdb *redis.Client) (*Game, error) {
	gameId := strconv.Itoa(int(b.Id))

	game := "game_content:" + gameId

	rdb.HSetNX(ctx, game, "Name", b.Name)
	rdb.HSetNX(ctx, game, "Description", b.Description)
	rdb.HSetNX(ctx, game, "Count", b.Count)
	rdb.Expire(ctx, game, 8*time.Hour)

	rdb.ZAdd(ctx, "game_index_cache:", &redis.Z{Score: float64(b.CreatedAt.Unix()), Member: game})
	rdb.Expire(ctx, "game_index_cache:", 8*time.Hour)

	gameData := rdb.HGetAll(ctx, game).Val()
	if len(gameData) == 0 {
		return nil, nil
	}

	i, _ := strconv.ParseInt(gameId, 10, 64)
	count, _ := strconv.ParseInt(gameData["Count"], 10, 64)

	return &Game{
		Id:          i,
		Name:        gameData["Name"],
		Description: gameData["Description"],
		Count:       count,
	}, nil
}
