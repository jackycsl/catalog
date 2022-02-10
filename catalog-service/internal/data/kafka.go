package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackycsl/catalog/catalog-service/internal/biz"
	"github.com/jackycsl/catalog/pkg/event/kafka"
	"github.com/jackycsl/catalog/pkg/util/helper"
)

func (r *gameRepo) KafkaCreateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	return nil, nil
}

func (r *gameRepo) KafkaUpdateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	return nil, nil
}

type GameEntry struct {
	GameId int64 `json:"game_id"`
}

func (r *gameRepo) KafkaBackfillGame(ctx context.Context, id int64) error {
	ge := &GameEntry{
		GameId: id,
	}
	g, err := json.Marshal(ge)
	if err != nil {
		return err
	}
	msg := kafka.NewMessage(helper.BackfillGameTopic, fmt.Sprintf("%d", id), g)
	err = r.data.kp.Send(context.Background(), msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Sending backfill message to Kafka, key:%s, value:%s\n", msg.Key(), msg.Value())
	return nil
}
