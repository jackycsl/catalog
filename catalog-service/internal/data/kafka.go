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

type GameEntry struct {
	GameId      int64  `json:"game_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Count       int64  `json:"count,omitempty"`
}

type GameList struct {
	PageNum  int64 `json:"page_num,omitempty"`
	PageSize int64 `json:"page_size,omitempty"`
}

func (r *gameRepo) KafkaCreateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	ge := &GameEntry{
		Name:        c.Name,
		Description: c.Description,
		Count:       c.Count,
	}
	g, err := json.Marshal(ge)
	if err != nil {
		return nil, err
	}
	msg := kafka.NewMessage(helper.CreateGameTopic, c.Name, g)

	go func() {
		fmt.Printf("Sending create message to Kafka, key:%s, value:%s\n", msg.Key(), msg.Value())
		err = r.data.kp.Send(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
	}()

	return &biz.Game{
		Name:        ge.Name,
		Description: ge.Description,
		Count:       ge.Count,
	}, nil
}

func (r *gameRepo) KafkaUpdateGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	return nil, nil
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

func (r *gameRepo) KafkaBackfillListGame(ctx context.Context, pageNum int64, pageSize int64) error {
	gl := &GameList{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
	g, err := json.Marshal(gl)
	if err != nil {
		return err
	}

	msg := kafka.NewMessage(helper.BackfillListGameTopic, fmt.Sprintf("%d-%d", pageNum, pageSize), g)
	err = r.data.kp.Send(context.Background(), msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Sending backfill List message to Kafka, key:%s, value:%s\n", msg.Key(), msg.Value())
	return nil
}
