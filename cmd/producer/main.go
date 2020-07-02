package producer

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/triardn/simple-kafka-producer-consumer/kafka"
)

var logger = log.With().Str("pkg", "main").Logger()

func main() {
	kafkaBrokerURLs := []string{"localhost:19092"}
	kafkaProducer, err := kafka.Init(kafkaBrokerURLs, "test-client", "topic-test-1")
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("unable to configure kafka")
		return
	}
	defer kafkaProducer.Close()

	msg := []byte("Test dong")

	err = kafka.Push(context.Background(), nil, msg)
	if err != nil {
		fmt.Println("error pushing data")

		return
	}

	fmt.Println("message successfully pushed")
}
