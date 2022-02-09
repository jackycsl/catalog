package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackycsl/catalog/pkg/event/event"
	"github.com/jackycsl/catalog/pkg/event/kafka"
)

func main() {
	sender, err := kafka.NewKafkaSender([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ {
		send(sender)
	}

	_ = sender.Close()
}

type GameEntry struct {
	GameId      int    `json:"game_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func send(sender event.Sender) {
	ge := &GameEntry{
		GameId:      1,
		Name:        "A Game",
		Description: "A new game",
	}

	g, err := json.Marshal(ge)
	if err != nil {
		fmt.Println(err)
	}

	msg := kafka.NewMessage("kratos", "kratos", g)
	err = sender.Send(context.Background(), msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key:%s, value:%s\n", msg.Key(), msg.Value())
}
