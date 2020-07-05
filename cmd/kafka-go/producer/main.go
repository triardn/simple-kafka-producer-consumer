package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/triardn/simple-kafka-producer-consumer/kafka"
)

var logger = log.With().Str("pkg", "main").Logger()

func main() {
	kafkaBrokerURLs := []string{"127.0.0.1:9092"}
	kafkaProducer, err := kafka.Init(kafkaBrokerURLs, "test-client", "topic-test-1")
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("unable to configure kafka")
		return
	}
	defer kafkaProducer.Close()

	for i := 0; i < 10; i++ {
		message := "Test Dong #" + string(i)

		msg := []byte(message)

		err = kafka.Push(context.Background(), nil, msg)
		if err != nil {
			fmt.Println("error pushing data")
			fmt.Println(err)

			return
		}

		fmt.Println("message successfully pushed")
	}
}
