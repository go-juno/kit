package main

import (
	"git.yupaopao.com/ops-public/kit/kafka"
	"context"
	"time"
)


func main() {
	ctx:=context.Background()
	newMqProducer := kafka.NewMqProducer(kafka.MqProducerOption{
		Addrs:     []string{"test-kafka.yupaopao.com:9092"},
		GroupId:   "public-kit-test",
		ClientId:  "public-kit-test",
		KeepOrder: false,
		Ctx:       ctx,
	})
	if err := newMqProducer.Start(); err != nil {
		panic(err)
	}
	defer newMqProducer.Stop()
	err := newMqProducer.SendMessageAsync("public_kit_kafka_test", "", []byte(time.Now().String()))
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Hour)
}

