package rocket

import (
	"testing"
)

func TestNewProducerRocketMq(t *testing.T) {
	c := &ProducerConfig{
		Addr: []string{"127.0.0.1:9876"},
		PrimitiveACL: PrimitiveACL{
			AccessKey: "wxx",
			SecretKey: "QWDARAED#RAAHGAT",
		},
		Retry: 1,
	}
	p, err := NewProducerRocketMq(c)
	if err != nil {
		return
	}
	err = p.ProducerStart()
	if err != nil {
		return
	}

	//err := p.SendAsync(context.Background(),)
	//if err != nil {
	//	return
	//}
	// return
}

func TestNewConsumeRocketMq(t *testing.T) {
	// return
}
