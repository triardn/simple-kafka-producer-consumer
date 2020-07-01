package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

// Push push message to kafka
func Push(ctx context.Context, key, value []byte) (err error) {
	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return writer.WriteMessages(ctx, msg)
}
