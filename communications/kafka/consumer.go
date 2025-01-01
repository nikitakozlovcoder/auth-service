package kafkaconsumer

import (
	"context"
	"errors"
	"io"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type ConsumerServer struct {
	readers []*kafka.Reader
}

type Consumer interface {
	Consume(even string, message []byte) error
}

type ConsumerConfig struct {
	GroupdID string
	Topic    string
}

func NewConsumer() *ConsumerServer {
	return &ConsumerServer{
		readers: make([]*kafka.Reader, 0),
	}
}

func (c *ConsumerServer) AddConsumer(cfg ConsumerConfig, consumer Consumer) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:  cfg.GroupdID,
		Topic:    cfg.Topic,
		MaxBytes: 10e6, // 10MB
	})

	c.readers = append(c.readers, r)

	go func() {
		for {
			ctx := context.Background()
			m, err := r.FetchMessage(ctx)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					slog.ErrorContext(ctx, err.Error(), slog.Any("error", err))
				}

				break
			}

			slog.InfoContext(ctx,
				"Consumed message",
				slog.String("topic", m.Topic),
				slog.Int("partition", m.Partition),
				slog.Int64("offset", m.Offset),
				slog.String("key", string(m.Key)),
				slog.String("value", string(m.Value)),
			)

			event := ""
			for _, h := range m.Headers {
				if h.Key == "event" {
					event = string(h.Value)
				}
			}

			err = consumer.Consume(event, m.Value)
			if err != nil {
				slog.ErrorContext(ctx, "failed to process message", slog.Any("error", err))
				continue
			}

			err = r.CommitMessages(ctx, m)
			if err != nil {
				slog.ErrorContext(ctx, "failed to commit messages", slog.Any("error", err))
			}
		}
	}()
}

func (c *ConsumerServer) Close() {
	for _, r := range c.readers {
		r.Close()
	}
}
