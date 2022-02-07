package data

import (
	"context"

	"entgo.io/ent/examples/fs/ent"
	"entgo.io/ent/examples/fs/ent/migrate"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jackycsl/catalog/catalog-job/internal/conf"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisClient, NewKafkaConsumer, NewCatalogJobRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
	kc  sarama.Consumer
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

func NewKafkaConsumer(conf *conf.Data, logger log.Logger) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}

// NewData
func NewData(entClient *ent.Client, redisClient *redis.Client, consumer sarama.Consumer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "catalog-job/data"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		kc:  consumer,
		log: log,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
		if err := d.kc.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
