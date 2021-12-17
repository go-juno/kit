package kafka

import (
	"context"

	"github.com/Shopify/sarama"
)

type MqConsumerOption struct {
	Addrs        []string
	ClientId     string
	GroupId      string
	Offset       ConsumerOffset // 默认取最新消息
	ProcessError bool
	Ctx          context.Context
	Handlers     map[string]MessageHandler
}

type MqProducerOption struct {
	Addrs     []string
	GroupId   string
	ClientId  string
	KeepOrder bool // 是否保证消息顺序
	Ctx       context.Context
}

type ConsumerOffset int

const (
	OffsetNewest = ConsumerOffset(sarama.OffsetNewest) // 最新消息
	OffsetOldest = ConsumerOffset(sarama.OffsetOldest) // 最老消息
)

type MessageHandler func(data []byte) bool
