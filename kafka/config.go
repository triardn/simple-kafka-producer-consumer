package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

// Init initialize kafka
func Init(kafkaBrokerURL []string, clientID string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}

	config := kafka.WriterConfig{
		Brokers:  kafkaBrokerURL,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
	}

	w = kafka.NewWriter(config)

	writer = w

	return
}
