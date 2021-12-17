package pulsar

import (
	"context"
	"fmt"
	"log"
	"testing"

	ps "github.com/apache/pulsar-client-go/pulsar"
)

func TestConsumer(t *testing.T) {
	client, err := NewClient(&Config{
		Host: "localhost",
		Port: 6650,
	})

	if err != nil {
		log.Println("err", err)
		return
	}
	var consumer *Consumer
	consumer, err = client.NewConsumer(ps.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "my-sub",
		Type:             ps.Shared,
	})

	if err != nil {
		log.Println("err", err)
		return
	}

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Println("err", err)
			return
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
	}

}

func TestProducer(t *testing.T) {
	client, err := NewClient(&Config{
		Host: "localhost",
		Port: 6650,
	})

	if err != nil {
		log.Println("err", err)
		return
	}
	var producer *Producer
	producer, err = client.NewProducer(ps.ProducerOptions{
		Topic: "my-topic",
	})

	if err != nil {
		log.Println("err", err)
		return
	}

	msg, err := producer.Send([]byte("dasdsad"))
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("msg", msg)
}
