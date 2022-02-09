package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackycsl/catalog/pkg/event/event"
	"github.com/jackycsl/catalog/pkg/event/kafka"
	"github.com/jackycsl/catalog/pkg/util/helper"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	receiver, err := kafka.NewKafkaReceiver([]string{"localhost:9092"}, helper.BackfillGameTopic)
	if err != nil {
		panic(err)
	}
	receive(receiver)

	<-sigs
	_ = receiver.Close()
}

type GameEntry struct {
	GameId int64 `json:"game_id"`
}

func receive(receiver event.Receiver) {
	fmt.Println("start receiver")
	err := receiver.Receive(context.Background(), func(ctx context.Context, msg event.Event) error {
		fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
		var game GameEntry
		if err := json.Unmarshal(msg.Value(), &game); err != nil {
			log.Println(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}
