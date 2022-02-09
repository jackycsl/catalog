package data

import (
	"context"

	"entgo.io/ent/examples/edgeindex/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jackycsl/catalog/catalog-service/internal/conf"
	"github.com/jackycsl/catalog/catalog-service/internal/data/ent"
	"github.com/jackycsl/catalog/pkg/event/event"
	"github.com/jackycsl/catalog/pkg/event/kafka"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisClient, NewKafkaProducer, NewGameRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
	kp  event.Sender
	log *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	//Run the auto migration tool
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false), migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewRedisClient(conf *conf.Data, logger log.Logger) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})

	return client
}

func NewKafkaProducer(conf *conf.Data, logger log.Logger) event.Sender {

	p, err := kafka.NewKafkaSender([]string(conf.Kafka.Addrs))
	if err != nil {
		panic(err)
	}
	return p
}

// NewData
func NewData(entClient *ent.Client, redisClient *redis.Client, producer event.Sender, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		kp:  producer,
		log: log,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
		if err := d.kp.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
