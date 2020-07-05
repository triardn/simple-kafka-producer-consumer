package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

func main() {
	cfg := kafka.ReaderConfig{
		Brokers:  []string{"127.0.0.1:9092"},
		GroupID:  "test-client",
		Topic:    "topic-test-1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	reader := kafka.NewReader(cfg)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error().Msgf("error while receiving message: %s", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
	}
}
