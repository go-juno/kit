使用开源的 [qlog](https://github.com/kkkbird/qlog)做一些改造

# 如何使用

## 基本使用

```go
package main

import (
 "git.yupaopao.com/ops-public/kit/log"
 "github.com/sirupsen/logrus"
 log1 "log"
)

var logger *logrus.Entry

func initLog() (err error) {
 entry, err := log.New("public-kit", "debug")
 if err != nil {
  return
 }
 logger = entry
 log.CollectSysLog() // 收集 sys log
 log1.Println("ok")
 log.Debug("log debug")
 return
}

```

## 自定义配置

```go

// 默认会读取当前 ".", "./conf",,"./etc" "./configs", "/etc/log" 目录下的 logger.yaml 做logger配置文件
//如果配置文件不满足该格式,调用该方法重新配置 logger

package main

import (
 "git.yupaopao.com/ops-public/kit/log"
 "github.com/sirupsen/logrus"
 log1 "log"
)

var logger *logrus.Entry

func initLog() (err error) {
 err = log.SetConfigFile("conf/app.yaml")
 if err != nil {
  panic(err)
 }
 entry := log.WithField("service","public-kit")
 logger = entry
 log.CollectSysLog() // 收集 sys log
 log1.Println("ok")
 log.Debug("log debug")
 return
}

// 或者在启动时指定 go run . --logger.config.file=./conf/app.yaml
```

## 原先已使用 logrus

log 劫持了 logrus 默认的 StandardLogger()，所以如果你的项目使用 logrus 没有 init logger 对象你可以只在你的主包中添加一行代码，并保持其他代码不变

```go
package main

import (
 _ "git.yupaopao.com/ops-public/kit/log"
 log "github.com/sirupsen/logrus"
)

func main() {
 log.Error("app error")
}
```

# 配置

支持以下 hooks:

- stdout
- stderr
- file

如果没有配置文件,默认会开启 stdout

```yaml
logger:
  level: debug # 日志等级 从低到高 trace|debug|info|warn|error|fatal|panic 默认 error
  reportcaller: false # 是否记录调用者信息,默认 false
  formatter: # 默认日志格式, stdout,file 可设置各自的 formatter
    name: classic # 默认 text,支持 text|classic|json|null
    opts:
      truncateCallerPath: true #
      callerPathStrip: true #
  stdout:
    enabled: true # 是否启用 stdout,默认禁用
    level: info # 日志等级 优先使用 logger.level,如 logger.level为 warn,stdout.level 只能为 warn 以上的等级

  file:
    enabled: true # 是否启用 file,默认禁用
    path: ./log/ # 日志文件路径
    name: message.log # 日志文件名称格式为 message-1635910826103494000.log
    level: trace
    formatter:
      name: json
      opts: # default formatter opts
        truncateCallerPath: true
        callerPathStrip: true
    rotate:
      time: 0 # 默认 24h,设置为 0 禁用 rotate
      maxage: 5m # 日志保留时间,默认 168h(7 days),设置 0 为禁止保留
      count: 0 # 文件保留数量,默认为0,禁止保留
```

# 格式

## classic 格式

```text
2021/11/03 11:49:08.113484+08:00 example/main.go:46 [W] hello 1 app=log event=1111 test=1
```

## text 格式

```text
WARN[0005]/Users/ltinyho/work/bx/ops/kit/log/example/main.go:46 main.ok() hello 1                                       app=log event=1111 test=1
```

## json 格式

```json
{
  "args": {
    "app": "log",
    "event": "1111",
    "test": "1"
  },
  "file": "example/main.go:46",
  "func": "main.ok",
  "level": "warning",
  "msg": "hello 1",
  "time": "2021/11/03 11:41:00.119193+08:00"
}
```
