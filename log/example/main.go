package main

import (
	"context"
	"time"

	"log"

	log1 "git.yupaopao.com/ops-public/kit/log"
)

func main() {
	time.Sleep(time.Second)
	// try to change the config now
	ctx, cancel := context.WithCancel(context.TODO())
	_, _ = log1.New("public-kit", "debug")
	log1.CollectSysLog()
	go func() {
		for i := 0; i < 100; i++ {
			log.Println(12321)
			time.Sleep(time.Second)
		}
		cancel()
	}()

	<-ctx.Done()
}
