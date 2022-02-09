package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/jackycsl/catalog/catalog-job/internal/conf"
	"github.com/jackycsl/catalog/catalog-job/internal/data"
	"github.com/jackycsl/catalog/pkg/event/kafka"
	"github.com/jackycsl/catalog/pkg/util/helper"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "game.job"
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
	defer db.Close()

	rdb := data.NewRedisClient(bc.Data)
	defer rdb.Close()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	receiver, err := kafka.NewKafkaReceiver(bc.Data.Kafka.Addrs, helper.BackfillGameTopic)
	if err != nil {
		panic(err)
	}

	data.Receive(receiver, db, rdb)

	<-sigs
	_ = receiver.Close()

}
