package rocket

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type ConsumerConfig struct {
	GroupName    string       `json:"group_name"`
	Addr         []string     `json:"addr"`
	PrimitiveACL PrimitiveACL `json:"primitive_acl"`
	BroadCast    bool         `json:"broad_cast"`
	Interceptor  bool         `json:"interceptor"`
	Namespace    string       `json:"namespace"`
	RetryOrder   bool         `json:"retry_order"`
	Strategy     bool         `json:"strategy"`
}

type ProducerConfig struct {
	Addr         []string     `json:"addr"`
	PrimitiveACL PrimitiveACL `json:"primitive_acl"`
	Retry        int          `json:"retry"`
	Namespace    string       `json:"namespace"`
	Interceptor  bool         `json:"interceptor"`
	TraceConfig  TraceConfig  `json:"trace_config"`
}
type TraceConfig struct {
	Access       int      `json:"access"`
	NamesrvAddrs []string `json:"namesrv_addrs"`
}
type PrimitiveACL struct {
	AccessKey string
	SecretKey string
}

type rocketPushConsumerService struct {
	PushConsumer rocketmq.PushConsumer
}
type rocketPullConsumerService struct {
	PullConsumer rocketmq.PullConsumer
}

type rocketProducerService struct {
	Producer            rocketmq.Producer
	TransactionProducer rocketmq.TransactionProducer
}

type MessageQueue struct {
	Topic      string `json:"topic"`
	BrokerName string `json:"brokerName"`
	QueueID    int    `json:"queue_id"`
}

type Message struct {
	Topic      string
	Body       []byte
	Properties map[string]string
}

type SendResult struct {
	SendResult *primitive.SendResult
}

type MessageExtResult struct {
	Result []*primitive.MessageExt
}

type PullResult struct {
	Result *primitive.PullResult
}
