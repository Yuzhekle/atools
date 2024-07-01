package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6388",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

func main() {
	go Publish()
	go Subscribe(1)
	go Subscribe(2)
	go Subscribe(3)
	select {}
}

func Publish() {
	for {
		rdb.Publish(context.Background(), "channel", "hello world")
		time.Sleep(5 * time.Second)
	}
}

func Subscribe(i int) {
	pubsub := rdb.Subscribe(context.Background(), "channel")

	for msg := range pubsub.Channel() {
		fmt.Printf("%s,%d,%s\n", msg.Channel, i, msg.Payload)
	}
}
