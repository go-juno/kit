package main

import (
	"context"
	"fmt"

	"github.com/go-juno/kit/kafka"
)

func main() {
	ctx := context.Background()
	consumer := kafka.NewMqConsumer(kafka.MqConsumerOption{
		Addrs:        []string{"127.0.0.1:9092"},
		ClientId:     "cloud-change",
		GroupId:      "cloud-change",
		Offset:       kafka.OffsetNewest,
		ProcessError: true,
		Ctx:          ctx,
		Handlers: map[string]kafka.MessageHandler{
			"public_kit_kafka_test": func(data []byte) bool {
				fmt.Println(string(data))
				return true
			},
		},
	})
	if err := consumer.Start(); err != nil {
		return
	}
	<-ctx.Done()
	consumer.Stop()
}
