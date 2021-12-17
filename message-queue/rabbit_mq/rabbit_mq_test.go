package rabbit

import "testing"

func TestNewRabbitClient(t *testing.T) {
	c := &MqConfig{
		"amqp://guest:guest@127.0.0.1:5672",
	}
	r, err := NewRabbitClient(c)
	if err != nil {
		panic(err)
	}
	r.channel.IsClosed()
}
