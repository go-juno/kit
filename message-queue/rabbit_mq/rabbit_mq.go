package rabbit

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/streadway/amqp"
	"golang.org/x/xerrors"
)

const delay = 3 // reconnect after delay seconds

// RabbitMQClient RabbitMQ结构体
type RabbitMQClient struct {
	conn    *Connection
	channel *Channel
}

// Channel amqp.Channel wapper
type Channel struct {
	*amqp.Channel
	closed int32

	CreateConsumeBeforeHook func() // 创建 consume 前 hook
}

// Connection amqp.Connection wrapper
type Connection struct {
	*amqp.Connection
	CreateChannelAfterHook func(ch *amqp.Channel) // 创建 channel 后 hook
}

func (c *Connection) ConnClose() bool {
	err := c.Close()
	return err == nil
}

// Channel wrap amqp.Connection.Channel, get a auto reconnect channel
func (c *Connection) Channel() (*Channel, error) {
	ch, err := c.Connection.Channel()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	if c.CreateChannelAfterHook != nil {
		c.CreateChannelAfterHook(ch)
	}
	channel := &Channel{
		Channel: ch,
	}
	go func() {
		for {
			reason, ok := <-channel.Channel.NotifyClose(make(chan *amqp.Error))
			// exit this goroutine if closed by developer
			if !ok || channel.IsClosed() {
				log.Printf("channel closed")
				channel.Close() // close again, ensure closed flag set when connection closed
				break
			}
			log.Printf("channel closed, reason: %v", reason)

			// reconnect if not closed by developer
			for {
				// wait 1s for connection reconnect
				time.Sleep(delay * time.Second)

				ch, err := c.Connection.Channel()
				if err == nil {
					log.Printf("channel recreate success")
					if c.CreateChannelAfterHook != nil {
						c.CreateChannelAfterHook(ch)
					}
					channel.Channel = ch
					break
				}
				log.Printf("channel recreate failed, err: %v", err)
			}
		}
	}()

	return channel, nil
}

// Dial wrap amqp.Dial, dial and get a reconnect connection
func Dial(url string) (*Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	connection := &Connection{
		Connection: conn,
	}

	go func() {
		for {
			reason, ok := <-connection.Connection.NotifyClose(make(chan *amqp.Error))
			// exit this goroutine if closed by developer
			if !ok {
				log.Printf("connection closed")
				break
			}
			log.Printf("connection closed, reason: %v", reason)

			// reconnect if not closed by developer
			for {
				// wait 1s for reconnect
				time.Sleep(delay * time.Second)

				conn, err := amqp.Dial(url)
				if err == nil {
					connection.Connection = conn
					log.Printf("reconnect success")
					break
				}

				log.Printf("reconnect failed, err: %v", err)
			}
		}
	}()
	return connection, nil
}

// IsClosed indicate closed by developer
func (ch *Channel) IsClosed() bool {
	return atomic.LoadInt32(&ch.closed) == 1
}

// Close ensure closed flag set
func (ch *Channel) Close() error {
	if ch.IsClosed() {
		return amqp.ErrClosed
	}

	atomic.StoreInt32(&ch.closed, 1)
	return ch.Channel.Close()
}

// Consume warp amqp.Channel.Consume, the returned delivery will end only when channel closed by developer
func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	deliveries := make(chan amqp.Delivery)

	go func() {
		for {
			if ch.CreateConsumeBeforeHook != nil {
				ch.CreateConsumeBeforeHook()
			}
			d, err := ch.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
			if err != nil {
				err = xerrors.Errorf("%w", err)
				log.Printf("consume failed, err: %v", err)
				time.Sleep(delay * time.Second)
				continue
			}

			for msg := range d {
				deliveries <- msg
			}

			// sleep before IsClose call. closed flag may not set before sleep.
			time.Sleep(delay * time.Second)
			if ch.IsClosed() {
				break
			}
		}
	}()

	return deliveries, nil
}

type MqConfig struct {
	MqURL string //amqp://guest:guest@127.0.0.1:5672
}

func NewRabbitClient(config *MqConfig) (RabbitMQClient, error) {
	rabbitConnection, err := Dial(config.MqURL)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return RabbitMQClient{}, err
	}

	rabbitMQChannel, err := rabbitConnection.Channel()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return RabbitMQClient{}, err
	}
	return RabbitMQClient{
		conn:    rabbitConnection,
		channel: rabbitMQChannel,
	}, nil
}
