package pulsar

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	"golang.org/x/xerrors"
)

type Config struct {
	Host string
	Port uint
}

type Client struct {
	client pulsar.Client
}

type Producer struct {
	producer pulsar.Producer
}

type Consumer struct {
	consumer pulsar.Consumer
}

func NewClient(config *Config) (client *Client, err error) {
	var c pulsar.Client
	c, err = pulsar.NewClient(pulsar.ClientOptions{
		URL: fmt.Sprintf("pulsar://%s:%d", config.Host, config.Port),
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	client = &Client{
		client: c,
	}
	return
}

func (c *Client) NewProducer(options pulsar.ProducerOptions) (producer *Producer, err error) {

	p, err := c.client.CreateProducer(options)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	producer = &Producer{
		producer: p,
	}
	return
}

func (p *Producer) Send(payload []byte) (message pulsar.MessageID, err error) {
	message, err = p.producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (c *Client) NewConsumer(options pulsar.ConsumerOptions) (consumer *Consumer, err error) {
	var cons pulsar.Consumer
	cons, err = c.client.Subscribe(options)

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	consumer = &Consumer{
		consumer: cons,
	}
	return
}

func (c *Consumer) Receive(ctx context.Context) (msg pulsar.Message, err error) {
	msg, err = c.consumer.Receive(context.Background())
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (c *Consumer) Ack(msg pulsar.Message) {
	c.consumer.Ack(msg)
}
