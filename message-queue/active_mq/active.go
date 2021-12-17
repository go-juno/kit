package active

import (
	"github.com/go-stomp/stomp/v3"
	"golang.org/x/xerrors"
)

type Config struct {
	Network     string
	DefaultPort int64
	ServerAddr  string

	Login Login
}

type Login struct {
	LoginName string
	PassCode  string
}

type MqClient interface {
	SendMessages(queueName string, contentType string, body []byte) error
	Subscribe(queueName string) (*Subscribe, error)
}

type mqClient struct {
	mqConn *stomp.Conn
}

type Subscribe struct {
	mqSubscribe *stomp.Subscription
}

func (m *mqClient) SendMessages(queueName string, contentType string, body []byte) error {
	err := m.mqConn.Send(queueName, contentType, body, nil)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func (m *mqClient) Subscribe(queueName string) (*Subscribe, error) {
	sub, err := m.mqConn.Subscribe(queueName, stomp.AckAuto)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return &Subscribe{
		mqSubscribe: sub,
	}, nil

}
func (s *Subscribe) RevMessages(messageCount int) [][]byte {
	result := make([][]byte, messageCount)
	for i := 1; i <= messageCount; i++ {
		msg := <-s.mqSubscribe.C
		actualText := msg.Body
		result = append(result, actualText)
	}
	return result
}

func NewActiveMQConn(config Config) (MqClient, error) {
	var options = []func(*stomp.Conn) error{
		stomp.ConnOpt.Login(config.Login.LoginName, config.Login.PassCode),
		stomp.ConnOpt.Host("/"),
	}
	conn, err := stomp.Dial(config.Network, config.ServerAddr, options...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &mqClient{
		mqConn: conn,
	}, nil
}
