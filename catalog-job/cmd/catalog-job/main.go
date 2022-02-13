package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/jackycsl/catalog/catalog-job/internal/conf"
	"github.com/jackycsl/catalog/catalog-job/internal/data"
	"github.com/jackycsl/catalog/pkg/event/kafka"
	"github.com/jackycsl/catalog/pkg/util/helper"
	"golang.org/x/sync/errgroup"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "game.catalog.job"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	db := data.NewEntClient(bc.Data)
	rdb := data.NewRedisClient(bc.Data)

	client, cleanup, err := data.NewData(db, rdb)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	app := data.NewGameRepo(client)

	g, ctx := errgroup.WithContext(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	backfillReceiver, err := kafka.NewKafkaReceiver(bc.Data.Kafka.Addrs, helper.BackfillGameTopic)
	if err != nil {
		panic(err)
	}

	backfillListReceiver, err := kafka.NewKafkaReceiver(bc.Data.Kafka.Addrs, helper.BackfillListGameTopic)
	if err != nil {
		panic(err)
	}

	createReceiver, err := kafka.NewKafkaReceiver(bc.Data.Kafka.Addrs, helper.CreateGameTopic)
	if err != nil {
		panic(err)
	}

	updateReceiver, err := kafka.NewKafkaReceiver(bc.Data.Kafka.Addrs, helper.UpdateGameTopic)
	if err != nil {
		panic(err)
	}

	g.Go(func() error {
		err = app.Backfill(backfillReceiver)
		if err != nil {
			log.Println(err)
		}
		return nil
	})

	g.Go(func() error {
		err = app.Create(createReceiver)
		if err != nil {
			log.Println(err)
		}
		return nil
	})

	g.Go(func() error {
		err = app.BackfillListGame(backfillListReceiver)
		if err != nil {
			log.Println(err)
		}
		return nil
	})

	g.Go(func() error {
		err = app.Update(updateReceiver)
		if err != nil {
			log.Println(err)
		}
		return nil
	})

	g.Go(func() error {
		select {
		case s := <-sigs:
			backfillReceiver.Close()
			createReceiver.Close()
			backfillListReceiver.Close()
			updateReceiver.Close()
			if err != nil {
				log.Println(err)
			}
			return fmt.Errorf("caught signal: %v", s.String())
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}

}
