package data

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/jackycsl/catalog/catalog-service/internal/biz"
	"github.com/jackycsl/catalog/pkg/util/helper"
)

func (r *gameRepo) KafkaCreateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	return nil, nil
}

func (r *gameRepo) KafkaUpdateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	return nil, nil
}

type GameEntry struct {
	GameId string `json:"game_id"`
}

func (r *gameRepo) KafkaBackfillGame(ctx context.Context, id int64) error {
	ge := &GameEntry{
		GameId: fmt.Sprintf("%d", id),
	}
	g, err := json.Marshal(ge)
	if err != nil {
		return err
	}
	r.data.kp.Input() <- &sarama.ProducerMessage{
		Topic: helper.BackfillGameTopic,
		Value: sarama.ByteEncoder(g),
	}
	return nil
}
