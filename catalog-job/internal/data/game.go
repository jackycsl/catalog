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

type GameList struct {
	PageNum  int64 `json:"page_num,omitempty"`
	PageSize int64 `json:"page_size,omitempty"`
}

func NewGameRepo(data *Data) biz.GameRepo {
	return &gameRepo{
		data: data,
	}
}

func (r *gameRepo) Backfill(receiver event.Receiver) error {
	var ctx = context.Background()
	fmt.Println("start receiver")

	err := receiver.Receive(ctx, func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var game biz.Game
		if err := json.Unmarshal(msg.Value(), &game); err != nil {
			log.Println(err)
		}
		ge, err := r.DbGetGame(ctx, game.Id)
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

func (r *gameRepo) Create(receiver event.Receiver) error {
	var ctx = context.Background()
	fmt.Println("start receiver")

	err := receiver.Receive(ctx, func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var game biz.Game
		if err := json.Unmarshal(msg.Value(), &game); err != nil {
			log.Println(err)
		}
		po, err := r.DbCreateGame(ctx, &game)
		if err != nil {
			return err
		}
		fmt.Printf("Game created at db, id: %d\n", po.Id)
		co, err := r.CacheCreateGame(ctx, po)
		if err != nil {
			return err
		}
		fmt.Printf("Game created at Cache key: %d\n", co.Id)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *gameRepo) BackfillListGame(receiver event.Receiver) error {
	var ctx = context.Background()
	fmt.Println("start receiver")

	err := receiver.Receive(ctx, func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var gamelist GameList
		if err := json.Unmarshal(msg.Value(), &gamelist); err != nil {
			log.Println(err)
		}
		// Pre heat
		count := gamelist.PageNum * gamelist.PageSize
		po, err := r.DbListGame(ctx, int64(1), count)
		if err != nil {
			return err
		}
		err = r.CacheSetGameList(ctx, po)
		if err != nil {
			return err
		}
		fmt.Printf("Games sets in cache\n")
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

func (r *gameRepo) Update(receiver event.Receiver) error {
	var ctx = context.Background()
	fmt.Println("start receiver")

	err := receiver.Receive(ctx, func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var game biz.Game
		if err := json.Unmarshal(msg.Value(), &game); err != nil {
			log.Println(err)
		}
		po, err := r.DbUpdateGame(ctx, &game)
		if err != nil {
			return err
		}
		fmt.Printf("Game updated at db, id: %d\n", po.Id)
		co, err := r.CacheCreateGame(ctx, po)
		if err != nil {
			return err
		}
		fmt.Printf("Game updated at Cache key: %d\n", co.Id)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
