package kafka

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
	"golang.org/x/xerrors"
)

var (
	queueSize      = 10000
	pendingErrSize = 10000
	closeReopen    = 0
	closeExit      = 1
)

type MqProducer interface {
	Start() error
	Stop()
	SendMessageAsync(topic, key string, message []byte) error // 异步发送消息
}

type mqProducer struct {
	addrs     []string
	clientId  string
	keepOrder bool
	producer  sarama.AsyncProducer
	queue     chan *sarama.ProducerMessage
	closeCh   chan int
	errMsg    chan *sarama.ProducerMessage
}

func NewMqProducer(opt MqProducerOption) MqProducer {
	producer := &mqProducer{
		addrs:     opt.Addrs,
		clientId:  opt.ClientId,
		keepOrder: opt.KeepOrder,
		queue:     make(chan *sarama.ProducerMessage, queueSize),
		producer:  nil,
		closeCh:   nil,
		errMsg:    make(chan *sarama.ProducerMessage, pendingErrSize),
	}
	return producer
}

func (p *mqProducer) newConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Net.MaxOpenRequests = 3
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Idempotent = false
	if p.keepOrder {
		config.Producer.Partitioner = sarama.NewReferenceHashPartitioner
	} else {
		config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	}
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = true
	config.Metadata.RefreshFrequency = time.Second * 5
	config.Version = sarama.V2_0_0_0
	if p.clientId != "" {
		config.ClientID = p.clientId
	}
	return config
}

func (p *mqProducer) createProducer() error {

	producer, err := sarama.NewAsyncProducer(p.addrs, p.newConfig())
	if err != nil {
		err = xerrors.Errorf("%w", err)
		log.Println(err)
		return err
	}
	p.producer = producer
	p.closeCh = make(chan int)
	return nil
}

func (p *mqProducer) Start() error {
	if err := p.createProducer(); err != nil {
		err = xerrors.Errorf("%w", err)
		log.Println("create producer error:", err)
		return err
	}
	go func() {
		for {
			if p.processSending() == closeExit {
				return
			}
			for err := p.createProducer(); err != nil; {
				err = xerrors.Errorf("%w", err)
				log.Println("create producer error:", err)
				time.Sleep(time.Second)
				continue
			}
		}
	}()

	return nil
}

func (p *mqProducer) Stop() {
	// send close cmd
	p.closeCh <- closeExit
	<-p.closeCh
	// wait the cmd to be processed
	if p.producer != nil {
		errs := p.producer.Close()
		emsgs, ok := errs.(sarama.ProducerErrors)
		if ok {
			for _, msg := range emsgs {
				p.storeErrMsg(msg.Msg)
			}
		}
	}
}

func (p *mqProducer) SendMessageAsync(topic, key string, message []byte) error {
	msg := &sarama.ProducerMessage{}
	msg.Value = sarama.ByteEncoder(message)
	msg.Topic = topic
	if key != "" {
		msg.Key = sarama.StringEncoder(key)
	}
	p.queue <- msg

	return nil
}

func (p *mqProducer) storeErrMsg(msg *sarama.ProducerMessage) {
	if msg == nil {
		return
	}
	select {
	case p.errMsg <- msg:
		return
	default:
		return
	}
}

func (p *mqProducer) processSending() int {
	defer func() {
		close(p.closeCh)
	}()
	for {
		select {
		case err := <-p.producer.Errors():
			if err != nil && err.Err != nil {
				p.processError(err.Err, err.Msg)
			}
		case msg := <-p.errMsg:
			p.producer.Input() <- msg
		case msg := <-p.queue:
			p.producer.Input() <- msg
		case code := <-p.closeCh:
			log.Println("receive close cmd, code=", code)
			return code
		}
	}
}

func (p *mqProducer) restart() {
	if p.closeCh == nil {
		return
	}
	go func() {
		p.closeCh <- closeReopen
	}()
}

func (p *mqProducer) processError(err error, msg *sarama.ProducerMessage) {
	p.storeErrMsg(msg)
	log.Println("err:", err)
	switch kerr := err.(type) {
	case sarama.KError:
		if kerr == sarama.ErrOutOfOrderSequenceNumber {
			log.Println("need to shutdown the producer and create an new one")
			// need to shut down the producer and create a new one
			p.restart()
		}
	}
}
