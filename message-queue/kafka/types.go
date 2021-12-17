package kafka

import (
	"context"
	"crypto/tls"
	kafkaMq "github.com/segmentio/kafka-go"
	"net"
	"time"
)

type Config struct {
	Network   string `json:"network"`
	Address   string `json:"address"`
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
	Dialer    Dialer `json:"dialer"`
}

// Connection kafka.Connection wrapper
type Connection struct {
	Conn *kafkaMq.Conn
}

type Dialer struct {
	ClientID string `json:"client_id"`

	// Optionally specifies the function that the dialer uses to establish
	// network connections. If nil, net.(*Dialer).DialContext is used instead.
	//
	// When DialFunc is set, LocalAddr, DualStack, FallbackDelay, and KeepAlive
	// are ignored.
	DialFunc func(ctx context.Context, network string, address string) (net.Conn, error)

	// Timeout is the maximum amount of time a dial will wait for a connect to
	// complete. If Deadline is also set, it may fail earlier.
	//
	// The default is no timeout.
	//
	// When dialing a name with multiple IP addresses, the timeout may be
	// divided between them.
	//
	// With or without a timeout, the operating system may impose its own
	// earlier timeout. For instance, TCP timeouts are often around 3 minutes.
	Timeout time.Duration `json:"timeout"`

	// Deadline is the absolute point in time after which dials will fail.
	// If Timeout is set, it may fail earlier.
	// Zero means no deadline, or dependent on the operating system as with the
	// Timeout option.
	Deadline time.Time `json:"deadline"`

	// LocalAddr is the local address to use when dialing an address.
	// The address must be of a compatible type for the network being dialed.
	// If nil, a local address is automatically chosen.
	LocalAddr net.Addr `json:"local_addr"`

	// DualStack enables RFC 6555-compliant "Happy Eyeballs" dialing when the
	// network is "tcp" and the destination is a host name with both IPv4 and
	// IPv6 addresses. This allows a client to tolerate networks where one
	// address family is silently broken.
	DualStack bool `json:"dual_stack"`

	// FallbackDelay specifies the length of time to wait before spawning a
	// fallback connection, when DualStack is enabled.
	// If zero, a default delay of 300ms is used.
	FallbackDelay time.Duration `json:"fallback_delay"`

	// KeepAlive specifies the keep-alive period for an active network
	// connection.
	// If zero, keep-alives are not enabled. Network protocols that do not
	// support keep-alives ignore this field.
	KeepAlive time.Duration `json:"keep_alive"`

	// TLS enables Dialer to open secure connections.  If nil, standard net.Conn
	// will be used.
	TLS *tls.Config `json:"tls"`

	SASLMechanism Mechanism `json:"sasl_mechanism"`
}

type ReaderConfig struct {
	// set offset excursion
	SetOffset            int64
	SetOffsetAtStartTime time.Time

	Brokers []string

	// GroupID holds the optional consumer group id.  If GroupID is specified, then
	// Partition should NOT be specified e.g. 0
	GroupID string

	// GroupTopics allows specifying multiple topics, but can only be used in
	// combination with GroupID, as it is a consumer-group feature. As such, if
	// GroupID is set, then either Topic or GroupTopics must be defined.
	GroupTopics []string

	// The topic to read messages from.
	Topic string

	// Partition to read messages from.  Either Partition or GroupID may
	// be assigned, but not both
	Partition int
	MinBytes  int

	// MaxBytes indicates to the broker the maximum batch size that the consumer
	// will accept. The broker will truncate a message to satisfy this maximum, so
	// choose a value that is high enough for your largest message size.
	//
	// Default: 1MB
	MaxBytes int
	// CommitInterval indicates the interval at which offsets are committed to
	// the broker.  If 0, commits will be handled synchronously.
	//
	// Default: 0
	//
	// Only used when GroupID is set
	CommitInterval time.Duration
}

type BalancerType string

const (
	BalanceHash       BalancerType = "Hash"
	BalanceLeastBytes BalancerType = "LeastBytes"
	BalanceCRC32B     BalancerType = "CRC32Balancer"
	BalanceRoundRobin BalancerType = "RoundRobin"
)

type TopicsConfig struct {

	// Topic name
	Topic string

	// NumPartitions created. -1 indicates unset.
	NumPartitions int

	// ReplicationFactor for the topic. -1 indicates unset.
	ReplicationFactor int
}

// SASLType  sasl support  type
type SASLType string

const (
	SASLPlain SASLType = "plain"
	SASLScram SASLType = "scram"
)

// Mechanism Plain
type Mechanism struct {
	Type     SASLType      `json:"type"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Timeout  time.Duration `json:"timeout"`
}

type WriterConfig struct {
	// This feild is required, attempting to write messages to a writer with a
	// nil address will error.
	Addr string

	// Topic is the name of the topic that the writer will produce messages to.
	//
	// Setting this field or not is a mutually exclusive option. If you set Topic
	// here, you must not set Topic for any produced Message. Otherwise, if you	do
	// not set Topic, every Message must have Topic specified.
	Topic string

	// The balancer used to distribute messages across partitions.
	//
	// The default is to use a round-robin distribution.
	Balancer BalancerType

	// Setting this flag to true causes the WriteMessages method to never block.
	// It also means that errors are ignored since the caller will not receive
	// the returned value. Use this only if you don't care about guarantees of
	// whether the messages were written to kafka.
	//
	// Defaults to false.sync send message or Async send message
	Async bool
	// Limit on how many attempts will be made to deliver a message.
	//
	// The default is to try at most 10 times.
	MaxAttempts int

	// Limit on how many messages will be buffered before being sent to a
	// partition.
	//
	// The default is to use a target batch size of 100 messages.
	BatchSize int

	// Limit the maximum size of a request in bytes before being sent to
	// a partition.
	//
	// The default is to use a kafka default value of 1048576.
	BatchBytes int64

	// Time limit on how often incomplete message batches will be flushed to
	// kafka.
	//
	// The default is to flush at least every second.
	BatchTimeout time.Duration

	// Timeout for read operations performed by the Writer.
	//
	// Defaults to 10 seconds.
	ReadTimeout time.Duration

	// Timeout for write operation performed by the Writer.
	//
	// Defaults to 10 seconds.
	WriteTimeout time.Duration
}

// Message is a data structure representing kafka messages.
type Message struct {
	// Topic indicates which topic this message was consumed from via Reader.
	//
	// When being used with Writer, this can be used to configured the topic if
	// not already specified on the writer itself.
	Topic string

	// Partition is read-only and MUST NOT be set when writing messages
	Partition     int
	Offset        int64
	HighWaterMark int64
	Key           []byte
	Value         []byte

	// If not set at the creation, Time will be automatically set when
	// writing the message.
	Time time.Time
}
