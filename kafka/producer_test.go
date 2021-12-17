package kafka

import (
	"context"
	"testing"
)

func TestNewMqProducer(t *testing.T) {
	newMqProducer := NewMqProducer(MqProducerOption{
		Addrs:     []string{"localhost:9092"},
		GroupId:   "cloud-change-test",
		ClientId:  "cloud-change-test",
		KeepOrder: false,
		Ctx:       context.Background(),
	})
	if err := newMqProducer.Start(); err != nil {
		t.Error(err)
	}
	defer newMqProducer.Stop()
	err := newMqProducer.SendMessageAsync("test", "", []byte("ok"))
	if err != nil {
		t.Error(err)
	}
}
