package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackycsl/catalog/catalog-job/internal/biz"
	"github.com/jackycsl/catalog/pkg/event/event"
)

var _ biz.GameRepo = (*gameRepo)(nil)

type gameRepo struct {
	data *Data
}

func NewGameRepo(data *Data) biz.GameRepo {
	return &gameRepo{
		data: data,
	}
}

type GameEntry struct {
	GameId int64 `json:"game_id"`
}

func (r *gameRepo) Backfill(receiver event.Receiver) error {
	var ctx = context.Background()
	fmt.Println("start receiver")

	err := receiver.Receive(ctx, func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var game GameEntry
		if err := json.Unmarshal(msg.Value(), &game); err != nil {
			log.Println(err)
		}
		ge, err := r.DbGetGame(ctx, game.GameId)
		if err != nil {
			return err
		}
		g, err := r.CacheSetGame(ctx, ge)
		if err != nil {
			return err
		}
		fmt.Printf("Cache Set, key: %d\n", g.Id)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
