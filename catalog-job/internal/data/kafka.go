package data

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"github.com/jackycsl/catalog/catalog-job/internal/biz"
)

func (r *catalogJobRepo) CacheSetGame(ctx context.Context, c *biz.Game) (*biz.Game, error) {
	return nil, nil
}

func (r *catalogJobRepo) KafkaConsumeBackfillGame(ctx context.Context) (*biz.Game, error) {
	partitionConsumer, err := r.data.kc.ConsumePartition("my_topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
		case <-partitionConsumer.Errors():
			break ConsumerLoop
		}
	}

	return nil, nil
}
