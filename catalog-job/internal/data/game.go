package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jackycsl/catalog/catalog-job/internal/data/ent"
	"github.com/jackycsl/catalog/pkg/event/event"
)

type GameEntry struct {
	GameId int64 `json:"game_id"`
}

func Receive(receiver event.Receiver, db *ent.Client, rdb *redis.Client) {
	var ctx = context.Background()
	fmt.Println("start receiver")
	id := int64(1)
	po, err := db.Game.Get(ctx, id)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(po)
	game := "game_content:1"
	gameData := rdb.HGetAll(ctx, game).Val()
	fmt.Println(gameData)
	err = receiver.Receive(ctx, func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var game GameEntry
		if err := json.Unmarshal(msg.Value(), &game); err != nil {
			log.Println(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}
