package kafka

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/IBM/sarama"
)

func TestMq(t *testing.T) {
	ctx := context.Background()
	client := NewMqConsumer(MqConsumerOption{
		Addrs:        []string{"localhost:9092"},
		ClientId:     "cloud-change-test",
		GroupId:      "cloud-change",
		Offset:       OffsetOldest,
		ProcessError: true,
		Ctx:          ctx,
		Handlers: map[string]MessageHandler{
			"test": func(data []byte) bool {
				fmt.Println("NewHandler", string(data))
				return true
			},
		},
	})
	if err := client.Start(); err != nil {
		t.Error(err)
		return
	}
	go producer()
	<-ctx.Done()
	client.Stop()
}

func producer() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = 1
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Metadata.RefreshFrequency = time.Second * 10
	syncProducer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	for range time.Tick(time.Second * 1) {
		val := sarama.StringEncoder(time.Now().Format("2016-01-02 15:04:05"))
		partition, offset, err := syncProducer.SendMessage(&sarama.ProducerMessage{
			Topic: "test",
			Key:   val,
			Value: val,
		})
		if err != nil {
			log.Println("producer err", err)

		} else {
			fmt.Printf("val:%s p:%d,o:%d\n", val, partition, offset)
		}
	}
}
