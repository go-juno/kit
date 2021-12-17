package sms

import (
	"log"
	"testing"
)

func TestSms(t *testing.T) {
	client, err := NewClient(&Config{
		AccessKeyId:     "LTAI5tN7h1N2iu2KvsfXnRcH",
		AccessKeySecret: "qFDnq48UGeLMUPU45p9R5WIcy1mUxe",
	})

	if err != nil {
		log.Println("err", err)
		return
	}

	err = client.SendSms("13122816381", "比心", "SMS_217895440", "daiasdsadasdsadaTest")
	if err != nil {
		log.Println("err", err)
		return
	}

}

func TestSmsCode(t *testing.T) {
	client, err := NewClient(&Config{
		AccessKeyId:     "LTAI5tN7h1N2iu2KvsfXnRcH",
		AccessKeySecret: "qFDnq48UGeLMUPU45p9R5WIcy1mUxe",
	})

	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("client", client)
	err = client.SendSms("18817814702", "比心", "SMS_217945442", "您的验证码为12312312")
	if err != nil {
		log.Println("err", err)
		return
	}

}
