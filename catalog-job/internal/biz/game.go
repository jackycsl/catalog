package biz

import (
	"context"
	"time"

	"github.com/jackycsl/catalog/pkg/event/event"
)

type Game struct {
	Id          int64     `json:"game_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Count       int64     `json:"count,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type GameRepo interface {
	Backfill(receiver event.Receiver) error
	DbGetGame(ctx context.Context, id int64) (*Game, error)
	CacheSetGame(ctx context.Context, g *Game) (*Game, error)

	Create(receiver event.Receiver) error
	DbCreateGame(ctx context.Context, g *Game) (*Game, error)

	BackfillListGame(receiver event.Receiver) error
	DBListGame(ctx context.Context, pageNum, pageSize int64) ([]*Game, error)
	CacheSetGameList(ctx context.Context, gs []*Game) error
}
