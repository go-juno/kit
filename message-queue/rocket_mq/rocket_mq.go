package rocket

import (
	"context"
	"fmt"
	"sync"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"golang.org/x/xerrors"
)

// PushConsumeService ConsumeService
type PushConsumeService interface {
	PushConsumerSubscribe(topic string) (*MessageExtResult, error)
	PushConsumerStart() error
	PushConsumerShutdown() error
	PushConsumerUnsubscribe(topic string) error
}

// PullConsumeService  ConsumeService
type PullConsumeService interface {
	PullFrom(ctx context.Context, queue MessageQueue, offset int64, numbers int) (*PullResult, error)
	Pull(ctx context.Context, topic string, numbers int) (*PullResult, error)
	Commit(ctx context.Context, mqs ...MessageQueue) (int64, error)
	CommittedOffset(queue MessageQueue) (int64, error)
	MessageQueues(topic string) []MessageQueue
	Seek(mq MessageQueue, offset int64) error
	Start() error
	Shutdown() error
}

// ProducerService  RocketMqService
type ProducerService interface {
	ProducerStart() error
	ProducerShutdown() error
	SendAsync(ctx context.Context, mgs []*Message) error
	SendOneWay(ctx context.Context, mq []*Message) error
	SendSync(ctx context.Context, mq []*Message) (*SendResult, error)
	TransactionSendMessageInTransaction(ctx context.Context, mq *Message) (*primitive.TransactionSendResult, error)
	Shutdown() error
}

func (r *rocketPushConsumerService) PushConsumerSubscribe(topic string) (*MessageExtResult, error) {
	var result []*primitive.MessageExt
	err := r.PushConsumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			result = append(result, msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return &MessageExtResult{Result: result}, nil
}

func (r *rocketPushConsumerService) PushConsumerStart() error {
	err := r.PushConsumer.Start()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func (r *rocketPushConsumerService) PushConsumerShutdown() error {
	err := r.PushConsumer.Shutdown()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil

}

func (r *rocketPushConsumerService) PushConsumerUnsubscribe(topic string) error {
	err := r.PushConsumer.Unsubscribe(topic)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func (r *rocketPullConsumerService) PullFrom(ctx context.Context, queue MessageQueue, offset int64, numbers int) (*PullResult, error) {
	mqQueue := primitive.MessageQueue{
		Topic:      queue.Topic,
		BrokerName: queue.BrokerName,
		QueueId:    queue.QueueID,
	}
	from, err := r.PullConsumer.PullFrom(ctx, mqQueue, offset, numbers)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return &PullResult{Result: from}, nil
}

func (r *rocketPullConsumerService) Pull(ctx context.Context, topic string, numbers int) (*PullResult, error) {
	from, err := r.PullConsumer.Pull(ctx, topic, numbers)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return &PullResult{Result: from}, nil
}

func (r *rocketPullConsumerService) Commit(ctx context.Context, mqs ...MessageQueue) (int64, error) {
	mQueues := make([]primitive.MessageQueue, len(mqs))
	for _, queue := range mqs {
		que := primitive.MessageQueue{
			Topic:      queue.Topic,
			BrokerName: queue.BrokerName,
			QueueId:    queue.QueueID,
		}
		mQueues = append(mQueues, que)
	}
	from, err := r.PullConsumer.Commit(ctx, mQueues...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return from, nil
}

func (r *rocketPullConsumerService) CommittedOffset(queue MessageQueue) (int64, error) {
	mqQueue := primitive.MessageQueue{
		Topic:      queue.Topic,
		BrokerName: queue.BrokerName,
		QueueId:    queue.QueueID,
	}
	from, err := r.PullConsumer.CommittedOffset(mqQueue)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return from, nil
}

func (r *rocketPullConsumerService) MessageQueues(topic string) []MessageQueue {
	results := make([]MessageQueue, 0)
	Queues := r.PullConsumer.MessageQueues(topic)
	for _, queue := range Queues {
		que := MessageQueue{
			Topic:      queue.Topic,
			BrokerName: queue.BrokerName,
			QueueID:    queue.QueueId,
		}
		results = append(results, que)
	}

	return results
}
func (r *rocketPullConsumerService) Seek(mq MessageQueue, offset int64) error {
	mqQueue := primitive.MessageQueue{
		Topic:      mq.Topic,
		BrokerName: mq.BrokerName,
		QueueId:    mq.QueueID,
	}
	err := r.PullConsumer.Seek(mqQueue, offset)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}
func (r *rocketPullConsumerService) Start() error {
	err := r.PullConsumer.Start()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}
func (r *rocketPullConsumerService) Shutdown() error {
	err := r.PullConsumer.Shutdown()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func (r *rocketProducerService) ProducerStart() error {
	err := r.Producer.Start()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}
func (r *rocketProducerService) ProducerShutdown() error {
	err := r.Producer.Shutdown()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func (r *rocketProducerService) SendAsync(ctx context.Context, mgs []*Message) error {
	if len(mgs) == 0 {
		return fmt.Errorf("send messages is none")
	}
	var wg sync.WaitGroup
	for _, m := range mgs {
		wg.Add(1)
		send := &primitive.Message{
			Topic: m.Topic,
			Body:  m.Body,
		}
		err := r.Producer.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, err error) {
			if err != nil {
				fmt.Printf("receive message error: %s\n", err)
			} else {
				fmt.Printf("send message success: result=%s\n", result.String())
			}
			wg.Done()
		}, send)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return err
		}
	}
	wg.Wait()

	return nil
}
func (r *rocketProducerService) SendOneWay(ctx context.Context, mq []*Message) error {
	mqs := make([]*primitive.Message, len(mq))
	for _, m := range mq {
		send := &primitive.Message{
			Topic: m.Topic,
			Body:  m.Body,
		}
		mqs = append(mqs, send)
	}
	err := r.Producer.SendOneWay(ctx, mqs...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}
func (r *rocketProducerService) SendSync(ctx context.Context, mq []*Message) (*SendResult, error) {
	mqs := make([]*primitive.Message, len(mq))
	for _, m := range mq {
		send := &primitive.Message{
			Topic: m.Topic,
			Body:  m.Body,
		}
		mqs = append(mqs, send)
	}
	sendResult, err := r.Producer.SendSync(ctx, mqs...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return &SendResult{SendResult: sendResult}, nil
}
func (r *rocketProducerService) TransactionSendMessageInTransaction(ctx context.Context, mq *Message) (*primitive.TransactionSendResult, error) {
	tranMessage := &primitive.Message{
		Topic: mq.Topic,
		Body:  mq.Body,
	}
	result, err := r.TransactionProducer.SendMessageInTransaction(ctx, tranMessage)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return result, nil
}
func (r *rocketProducerService) Shutdown() error {
	err := r.Producer.Shutdown()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func NewPushConsumeRocketMq(conf *ConsumerConfig) (PushConsumeService, error) {
	option := make([]consumer.Option, 0)
	option = append(option, consumer.WithGroupName(conf.GroupName))
	option = append(option, consumer.WithNsResolver(primitive.NewPassthroughResolver(conf.Addr)))
	if conf.PrimitiveACL.AccessKey != "" && conf.PrimitiveACL.SecretKey != "" {
		option = append(option, consumer.WithCredentials(primitive.Credentials{
			AccessKey: conf.PrimitiveACL.AccessKey,
			SecretKey: conf.PrimitiveACL.SecretKey,
		}))
	}
	if conf.BroadCast {
		option = append(option, consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset))
		option = append(option, consumer.WithConsumerModel(consumer.BroadCasting))
	}
	if conf.Interceptor {
		option = append(option, consumer.WithInterceptor(UserFirstInterceptor(), UserSecondInterceptor()))
	}
	if conf.Namespace != "" {
		option = append(option, consumer.WithNamespace(conf.Namespace))
	}
	if conf.RetryOrder {
		option = append(option, consumer.WithConsumerOrder(true))
		option = append(option, consumer.WithMaxReconsumeTimes(5))
	}
	if conf.Strategy {
		option = append(option, consumer.WithStrategy(consumer.AllocateByAveragely))
	}
	pushConsumer, err := rocketmq.NewPushConsumer(option...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		fmt.Printf("init consumer error: " + err.Error())
		return nil, err
	}
	return &rocketPushConsumerService{
		PushConsumer: pushConsumer,
	}, nil
}

func NewPllConsumeRocketMq(conf *ConsumerConfig) (PullConsumeService, error) {
	option := make([]consumer.Option, 0)
	option = append(option, consumer.WithGroupName(conf.GroupName))
	option = append(option, consumer.WithNsResolver(primitive.NewPassthroughResolver(conf.Addr)))
	pullConsumer, err := rocketmq.NewPullConsumer(option...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		fmt.Printf("init consumer error: " + err.Error())
		return nil, err
	}
	return &rocketPullConsumerService{PullConsumer: pullConsumer}, nil
}

func NewProducerRocketMq(conf *ProducerConfig) (ProducerService, error) {
	option := make([]producer.Option, 0)
	option = append(option, producer.WithNsResolver(primitive.NewPassthroughResolver(conf.Addr)))
	option = append(option, producer.WithRetry(conf.Retry))
	option = append(option, producer.WithQueueSelector(producer.NewManualQueueSelector()))
	if conf.PrimitiveACL.AccessKey != "" && conf.PrimitiveACL.SecretKey != "" {
		option = append(option, producer.WithCredentials(primitive.Credentials{
			AccessKey: conf.PrimitiveACL.AccessKey,
			SecretKey: conf.PrimitiveACL.SecretKey,
		}))
	}
	if conf.Interceptor {
		option = append(option, producer.WithInterceptor(UserFirstInterceptor(), UserSecondInterceptor()))
	}
	if conf.Namespace != "" {
		option = append(option, producer.WithNamespace(conf.Namespace))
	}
	if len(conf.TraceConfig.NamesrvAddrs) > 0 && conf.TraceConfig.Access > 0 {
		traceCfg := &primitive.TraceConfig{
			Access:       primitive.AccessChannel(conf.TraceConfig.Access),
			NamesrvAddrs: conf.TraceConfig.NamesrvAddrs,
		}
		option = append(option, producer.WithTrace(traceCfg))
	}

	newProducer, err := rocketmq.NewProducer(option...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &rocketProducerService{
		Producer: newProducer,
	}, nil
}

func UserFirstInterceptor() primitive.Interceptor {
	return func(ctx context.Context, req, reply interface{}, next primitive.Invoker) error {
		fmt.Printf("user first interceptor before invoke: req:%v\n", req)
		err := next(ctx, req, reply)
		fmt.Printf("user first interceptor after invoke: req: %v, reply: %v \n", req, reply)
		return err
	}
}

func UserSecondInterceptor() primitive.Interceptor {
	return func(ctx context.Context, req, reply interface{}, next primitive.Invoker) error {
		fmt.Printf("user second interceptor before invoke: req: %v\n", req)
		err := next(ctx, req, reply)
		fmt.Printf("user second interceptor after invoke: req: %v, reply: %v \n", req, reply)
		return err
	}
}
