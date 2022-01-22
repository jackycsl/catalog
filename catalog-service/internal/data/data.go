package data

import (
	"context"

	"entgo.io/ent/examples/edgeindex/ent/migrate"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jackycsl/catalog/catalog-service/internal/conf"
	"github.com/jackycsl/catalog/catalog-service/internal/data/ent"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisClient, NewKafkaProducer, NewGameRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
	kp  sarama.AsyncProducer
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
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
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

func NewKafkaProducer(conf *conf.Data, logger log.Logger) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}

// NewData .
func NewData(entClient *ent.Client, redisClient *redis.Client, producer sarama.AsyncProducer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
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
