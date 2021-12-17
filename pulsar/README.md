# Pulsar

## 目前提供 consumer 和 newProducer

## Usage

```go
    go get -u git.yupaopao.com/ops-public/kit/pulsar
    import "git.yupaopao.com/ops-public/kit/pulsar"

```

### 注意

```text
    由于pulsar官方包引用了一个包不再升级, 有可能出现 kcItem.SetAccess undefined 的错误
    解决方案:
    go.mod 最后一行添加
    replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
    再重新运行即可

```

### consumer 示例

```go
package main

import (
 "context"
 "fmt"
 "log"

 "git.yupaopao.com/ops-public/kit/pulsar"
 ps "github.com/apache/pulsar-client-go/pulsar"
)

func main() {
 client, err := pulsar.NewClient(&pulsar.Config{
  Host: "localhost",
  Port: 6650,
 })

 if err != nil {
  log.Println("err", err)
  return
 }
 consumer, err := client.NewConsumer(ps.ConsumerOptions{
  Topic:            "my-topic",
  SubscriptionName: "my-sub",
  Type:             ps.Shared,
 })

 for {
  msg, err := consumer.Receive(context.Background())
  if err != nil {
   log.Println("err", err)
   return
  }

  fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
   msg.ID(), string(msg.Payload()))

  consumer.Ack(msg)
 }

}

```

### newProducer 示例

```go
package main

import (
 "log"

 "git.yupaopao.com/ops-public/kit/pulsar"
 ps "github.com/apache/pulsar-client-go/pulsar"
)

func main() {
 client, err := pulsar.NewClient(&pulsar.Config{
  Host: "localhost",
  Port: 6650,
 })

 if err != nil {
  log.Println("err", err)
  return
 }
 newProducer, err := client.NewProducer(ps.ProducerOptions{
  Topic: "my-topic",
 })

 msg, err := newProducer.Send([]byte("dasdsad"))
 if err != nil {
  log.Println("err", err)
  return
 }
 log.Println("msg", msg)
}

```
