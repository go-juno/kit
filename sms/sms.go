package sms

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"golang.org/x/xerrors"
)

type Config struct {
	AccessKeyId     string
	AccessKeySecret string
}

type Client struct {
	client *dysmsapi.Client
}

func NewClient(config *Config) (client *Client, err error) {

	c := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &config.AccessKeyId,
		// 您的AccessKey Secre
		AccessKeySecret: &config.AccessKeySecret,
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	smsClient, err := dysmsapi.NewClient(c)

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	client = &Client{
		client: smsClient,
	}
	return
}

func (c *Client) SendSms(phone, signature, templateCode, content string) (err error) {
	request := &dysmsapi.SendSmsRequest{
		SignName:      &signature,
		PhoneNumbers:  &phone,
		TemplateCode:  &templateCode,
		TemplateParam: tea.String(fmt.Sprintf("{\"content\":\"%s\n\"}", content)),
	}
	_, err = c.client.SendSms(request)

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
