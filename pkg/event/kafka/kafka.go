package kafka

import (
	"context"
	"log"

	"github.com/jackycsl/catalog/pkg/event/event"
	"github.com/segmentio/kafka-go"
)

var (
	_ event.Sender   = (*kafkaSender)(nil)
	_ event.Receiver = (*kafkaReceiver)(nil)
	_ event.Event    = (*Message)(nil)
)

type Message struct {
	topic string
	key   string
	value []byte
}

func (m *Message) Topic() string {
	return m.topic
}

func (m *Message) Key() string {
	return m.key
}

func (m *Message) Value() []byte {
	return m.value
}

func NewMessage(topic string, key string, value []byte) event.Event {
	return &Message{
		topic: topic,
		key:   key,
		value: value,
	}
}

type kafkaSender struct {
	writer *kafka.Writer
	// topic  string
}

func (s *kafkaSender) Send(ctx context.Context, message event.Event) error {
	err := s.writer.WriteMessages(ctx, kafka.Message{
		Topic: message.Topic(),
		Key:   []byte(message.Key()),
		Value: message.Value(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *kafkaSender) Close() error {
	err := s.writer.Close()
	if err != nil {
		return nil
	}
	return nil
}

func NewKafkaSender(address []string) (event.Sender, error) {
	w := &kafka.Writer{
		// Topic:    topic,
		Addr:     kafka.TCP(address...),
		Balancer: &kafka.LeastBytes{},
	}
	return &kafkaSender{writer: w}, nil
}

type kafkaReceiver struct {
	reader *kafka.Reader
	topic  string
}

func (k *kafkaReceiver) Receive(ctx context.Context, handler event.Handler) error {
	go func() {
		for {
			m, err := k.reader.FetchMessage(context.Background())
			if err != nil {
				break
			}
			err = handler(context.Background(), &Message{
				key:   string(m.Key),
				value: m.Value,
			})
			if err != nil {
				log.Fatal("message handling exception:", err)
			}
			if err := k.reader.CommitMessages(ctx, m); err != nil {
				log.Fatal("failed to commit messages:", err)
			}
		}
	}()
	return nil
}

func (k *kafkaReceiver) Close() error {
	err := k.reader.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewKafkaReceiver(address []string, topic string) (event.Receiver, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  address,
		GroupID:  "group-a",
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &kafkaReceiver{reader: r, topic: topic}, nil
}
