package kafka

import (
	"context"
	"errors"
	"log"

	"github.com/IBM/sarama"
	"golang.org/x/xerrors"
)

var (
	ConsumerBrokerNotFound = errors.New("can not found the Broker address")
)

type MqConsumer interface {
	Start() error
	Stop()
}

type mqConsumer struct {
	addrs        []string
	clientId     string
	groupId      string
	topics       []string
	offset       ConsumerOffset // 默认取最新消息
	processError bool
	ctx          context.Context
	consumer     sarama.ConsumerGroup
	done         chan struct{}
	handlers     map[string]MessageHandler
}

func NewMqConsumer(opt MqConsumerOption) MqConsumer {
	consumer := &mqConsumer{
		addrs:        opt.Addrs,
		clientId:     opt.ClientId,
		groupId:      opt.GroupId,
		handlers:     opt.Handlers,
		offset:       opt.Offset,
		processError: opt.ProcessError,
		ctx:          opt.Ctx,
		done:         make(chan struct{}, 1),
		consumer:     nil,
	}
	topics := make([]string, 0, len(consumer.handlers))
	for k := range opt.Handlers {
		topics = append(topics, k)
	}
	consumer.topics = topics
	return consumer
}

func (c *mqConsumer) Start() error {
	err := c.create()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	if c.ctx == nil {
		c.ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(c.ctx)
	c.ctx = ctx
	if c.processError {
		go c.doConsumerErrors()
	}
	go func() {
		for {
			select {
			case <-c.done:
				cancel()
				return
			case <-c.ctx.Done():
				return
			default:
				err = c.consumer.Consume(c.ctx, c.topics, c)
				if err != nil {
					err = xerrors.Errorf("%w", err)
					log.Println("consume err:", err)
				}
			}
		}
	}()
	return nil
}

func (c *mqConsumer) Stop() {
	c.done <- struct{}{}
	if c.consumer != nil {
		c.consumer.Close()
		c.consumer = nil
	}
	log.Println("consumer: Stop")
}

func (c *mqConsumer) doConsumerErrors() {
	for err := range c.consumer.Errors() {
		if err == nil {
			log.Println("errors return nil")
		}
		log.Println("consumer ERROR:", err)
	}
	log.Println("doConsumerErrors stop")
}

func (c *mqConsumer) create() error {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0

	offset := sarama.OffsetNewest
	if c.offset == OffsetOldest {
		offset = sarama.OffsetOldest
	}
	config.Consumer.Offsets.Initial = offset

	if c.clientId != "" {
		config.ClientID = c.clientId
	}
	if len(c.addrs) == 0 {
		return ConsumerBrokerNotFound
	}
	group, err := sarama.NewConsumerGroup(c.addrs, c.groupId, config)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	c.consumer = group
	return nil
}

func (c *mqConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if handler, ok := c.handlers[msg.Topic]; ok {
			handler(msg.Value)
			sess.MarkMessage(msg, "")
		} else {
			log.Println("Unexpected topic ", msg.Topic)
		}
	}
	return nil
}

func (mqConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	log.Println("consumer Setup")
	return nil
}

func (mqConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Println("consumer Cleanup")
	return nil
}
