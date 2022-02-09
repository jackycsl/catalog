package data

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jackycsl/catalog/catalog-job/internal/conf"
	"github.com/jackycsl/catalog/catalog-job/internal/data/ent"
	"github.com/jackycsl/catalog/catalog-job/internal/data/ent/migrate"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func NewEntClient(conf *conf.Data) *ent.Client {

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewRedisClient(conf *conf.Data) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})

	return client
}
