package biz

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackycsl/catalog/pkg/util/helper"
	"golang.org/x/sync/singleflight"
)

type Image struct {
	URL string
}

type Game struct {
	Id          int64
	Name        string
	Description string
	Count       int64
}

type GameRepo interface {
	// db
	// read from db if cache miss
	GetGame(ctx context.Context, id int64) (*Game, error)
	ListGame(ctx context.Context, pageNum, pageSize int64) ([]*Game, error)
	CreateGame(ctx context.Context, c *Game) (*Game, error)
	UpdateGame(ctx context.Context, c *Game) (*Game, error)

	// redis
	// read from Redis. if cache miss, read from db and create backfill job
	CacheGetGame(ctx context.Context, id int64) (*Game, error)
	CacheListGame(ctx context.Context, pageNum int64, pageSize int64) ([]*Game, error)
	CacheCreateGame(ctx context.Context, c *Game) (*Game, error)
	CacheUpdateGame(ctx context.Context, c *Game) (*Game, error)

	// kafka
	// Create, update & backfill will be handled by catalog-job
	KafkaCreateGame(ctx context.Context, c *Game) (*Game, error)
	KafkaUpdateGame(ctx context.Context, c *Game) (*Game, error)
	KafkaBackfillGame(ctx context.Context, Id int64) error
}

type GameUseCase struct {
	repo GameRepo
	log  *log.Helper
}

func NewGameUseCase(repo GameRepo, logger log.Logger) *GameUseCase {
	return &GameUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/game"))}
}

func (uc *GameUseCase) Create(ctx context.Context, u *Game) (*Game, error) {
	g, err := uc.repo.KafkaCreateGame(ctx, u)
	if err != nil {
		return nil, err
	}
	return g, err
}

func (uc *GameUseCase) Get(ctx context.Context, id int64) (*Game, error) {
	var requestGroup singleflight.Group

	p, err := uc.repo.CacheGetGame(ctx, id)

	// cache miss, read from db, create backfill job
	if errors.Is(err, helper.ErrRecordNotFound) {
		val, err, _ := requestGroup.Do("get from db", func() (interface{}, error) {
			fmt.Println("requesting from db")
			return uc.repo.GetGame(ctx, id)
		})
		if err != nil {
			return nil, err
		}
		p = val.(*Game)

		go func(ctx context.Context, id int64) {
			uc.repo.KafkaBackfillGame(ctx, id)
		}(ctx, id)
	}

	return p, nil
}

func (uc *GameUseCase) Update(ctx context.Context, u *Game) (*Game, error) {
	return uc.repo.UpdateGame(ctx, u)
}

func (uc *GameUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Game, error) {
	return uc.repo.ListGame(ctx, pageNum, pageSize)
}
