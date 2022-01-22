package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
)

type Image struct {
	URL string
}

type Game struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	Images      []Image
}

type GameRepo interface {
	// db
	// read from db if cache miss
	GetGame(ctx context.Context, id int64) (*Game, error)
	ListGame(ctx context.Context, pageNum, pageSize int64) ([]*Game, error)
	// CreateGame(ctx context.Context, c *Game) (*Game, error)
	// UpdateGame(ctx context.Context, c *Game) (*Game, error)

	// redis
	// read from Redis. if cache miss, read from db and create backfill job
	CacheGetGame(ctx context.Context, id int64) (*Game, error)
	CacheListGame(ctx context.Context, pageNum int64, pageSize int64) ([]*Game, error)

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
	return uc.repo.KafkaCreateGame(ctx, u)
}

func (uc *GameUseCase) Get(ctx context.Context, id int64) (*Game, error) {

	p, err := uc.repo.CacheGetGame(ctx, id)
	if err != nil {
		return nil, err
	}
	g, ctx := errgroup.WithContext(ctx)

	// cache miss, read from db, create backfill job
	if p.Id == 0 {
		p, err = uc.repo.GetGame(ctx, id)
		if err != nil {
			return nil, err
		}

		g.Go(func() error {
			err = uc.repo.KafkaBackfillGame(ctx, id)
			return err
		})

	}

	return p, nil
}

func (uc *GameUseCase) Update(ctx context.Context, u *Game) (*Game, error) {
	return uc.repo.KafkaUpdateGame(ctx, u)
}

func (uc *GameUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Game, error) {
	return uc.repo.ListGame(ctx, pageNum, pageSize)
}
