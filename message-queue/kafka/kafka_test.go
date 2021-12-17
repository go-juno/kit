package kafka

import (
	"fmt"
	"testing"
	"time"
)

func TestNewKafkaConn(t *testing.T) {
	// var ti time.Time
	//fmt.Println(ti.IsZero())
	c, err := NewKafkaConn(&Config{
		Network:   "",
		Address:   "",
		Topic:     "",
		Partition: 0,
		Dialer:    Dialer{},
	})
	if err != nil {
		return
	}
	reader, err := c.Reader(ReaderConfig{
		SetOffset:            0,
		SetOffsetAtStartTime: time.Time{},
		Brokers:              nil,
		GroupID:              "",
		GroupTopics:          nil,
		Topic:                "",
		Partition:            0,
		MinBytes:             0,
		MaxBytes:             0,
		CommitInterval:       0,
	})
	if err != nil {
		return
	}
	message, err := reader.ReadMessage(54)
	if err != nil {
		return
	}
	fmt.Println(message)
}
