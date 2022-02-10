package biz

import (
	"context"
	"time"

	"github.com/jackycsl/catalog/pkg/event/event"
)

type Game struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	CreatedAt   time.Time
}

type GameRepo interface {
	Backfill(receiver event.Receiver) error
	DbGetGame(ctx context.Context, id int64) (*Game, error)
	CacheSetGame(ctx context.Context, b *Game) (*Game, error)
}
