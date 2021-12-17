package kafka

import (
	"context"
	"net"
	"strconv"
	"time"

	kafkaMq "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"golang.org/x/xerrors"
)

type Writer struct {
	*kafkaMq.Writer
}

type Reader struct {
	*kafkaMq.Reader
}

// NewKafkaConn kafka To Connect To Leader Via a Non-leader Connection
func NewKafkaConn(config *Config) (*Connection, error) {
	var err error
	var conn *kafkaMq.Conn
	if config.Dialer.DualStack {
		dialer := &kafkaMq.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
			TLS:       config.Dialer.TLS,
		}
		if config.Dialer.SASLMechanism.Type != "" {
			if config.Dialer.SASLMechanism.Type == SASLScram {
				scramMechanism, err := scram.Mechanism(scram.SHA512, config.Dialer.SASLMechanism.Username, config.Dialer.SASLMechanism.Password)
				if err != nil {
					err = xerrors.Errorf("%w", err)
					return nil, err
				}
				dialer.SASLMechanism = scramMechanism
			} else {
				plainMechanism := plain.Mechanism{
					Username: config.Dialer.SASLMechanism.Username,
					Password: config.Dialer.SASLMechanism.Password,
				}
				dialer.SASLMechanism = plainMechanism
			}
		}
		conn, err = dialer.DialContext(context.Background(), config.Network, config.Address)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return nil, err
		}
	} else {
		conn, err = kafkaMq.DialLeader(context.Background(), config.Network, config.Address, config.Topic, config.Partition)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return nil, err
		}
	}
	return &Connection{Conn: conn}, nil
} // LeaderConn  requests kafka for the current controller and returns its URL
func (c *Connection) LeaderConn() (*Connection, error) {
	controller, err := c.Conn.Controller()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	var connLeader *kafkaMq.Conn
	connLeader, err = kafkaMq.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return &Connection{Conn: connLeader}, nil
} // CloseConn 关闭链接
func (c *Connection) CloseConn() error {
	err := c.Conn.Close()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
} // DeleteTopics deletes the specified topics.
func (c *Connection) DeleteTopics(topics ...string) error {
	err := c.Conn.DeleteTopics(topics...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

// CreatTopics  creates one topic per provided configuration with idempotent
// operational semantics. In other words, if CreateTopics is invoked with a
// configuration for an existing topic, it will have no effect.
func (c *Connection) CreatTopics(topicsConfig *TopicsConfig) error {
	// to create topics when auto.create.topics.enable='true'
	topicConfigs := []kafkaMq.TopicConfig{
		{
			Topic:             topicsConfig.Topic,
			NumPartitions:     topicsConfig.NumPartitions,
			ReplicationFactor: topicsConfig.ReplicationFactor,
		},
	}
	err := c.Conn.CreateTopics(topicConfigs...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
} // TopicsList To list topics
func (c *Connection) TopicsList() (map[string]struct{}, error) {
	partitions, err := c.Conn.ReadPartitions()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	m := map[string]struct{}{}
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	return m, nil
} // Reader provides a high-level API for consuming messages from kafka.
func (c *Connection) Reader(config ReaderConfig) (*Reader, error) {
	reader := kafkaMq.NewReader(kafkaMq.ReaderConfig{
		Brokers:        config.Brokers,
		Topic:          config.Topic,
		GroupID:        config.GroupID,
		Partition:      config.Partition,
		MinBytes:       config.MinBytes, // 10KB
		MaxBytes:       config.MaxBytes, // 10MB
		CommitInterval: config.CommitInterval,
	})
	if config.SetOffset > 0 {
		err := reader.SetOffset(config.SetOffset)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return nil, err
		}
	}
	if !config.SetOffsetAtStartTime.IsZero() {
		err := reader.SetOffsetAt(context.Background(), config.SetOffsetAtStartTime)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return nil, err
		}
	}
	return &Reader{
		reader,
	}, nil
} // Writer Methods of Writer are safe to use concurrently from multiple goroutines,
func (c *Connection) Writer(config WriterConfig) *Writer {
	writer := &kafkaMq.Writer{
		Addr:     kafkaMq.TCP(config.Addr),
		Topic:    config.Topic,
		Balancer: Balance(config.Balancer),
		Async:    config.Async,
	}
	return &Writer{writer}
}
func Balance(balance BalancerType) kafkaMq.Balancer {
	if balance == "Hash" {
		return &kafkaMq.Hash{}
	}
	if balance == "LeastBytes" {
		return &kafkaMq.LeastBytes{}
	}
	if balance == "CRC32Balancer" {
		return &kafkaMq.CRC32Balancer{}
	}
	if balance == "RoundRobin" {
		return &kafkaMq.RoundRobin{}
	}
	if balance == "" {
		return nil
	}
	return nil
}
func (r *Reader) ReadMessage(messageCount int) ([]Message, error) {
	results := make([]Message, messageCount)
	for i := 1; i < messageCount; i++ {
		km, err := r.Reader.ReadMessage(context.Background())
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return nil, err
		}
		m := Message{
			Topic:         km.Topic,
			Partition:     km.Partition,
			Offset:        km.Offset,
			HighWaterMark: km.HighWaterMark,
			Key:           km.Key,
			Value:         km.Value,
			Time:          km.Time,
		}
		results = append(results, m)
	}
	return results, nil
}

func (r *Writer) KafKaClose() error {
	err := r.Writer.Close()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func (r *Writer) WriteMessage(messAgs []Message) error {
	results := make([]kafkaMq.Message, len(messAgs))
	for _, m := range messAgs {
		km := kafkaMq.Message{
			Topic:         m.Topic,
			Partition:     m.Partition,
			Offset:        m.Offset,
			HighWaterMark: m.HighWaterMark,
			Key:           m.Key,
			Value:         m.Value,
		}
		results = append(results, km)
	}
	err := r.Writer.WriteMessages(context.Background(), results...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}
