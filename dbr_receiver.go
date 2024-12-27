package main

import (
	"fmt"
	"log"
	"time"
)

// MyEventReceiver 实现 dbr.EventReceiver 接口
type MyEventReceiver struct {
	logger *log.Logger
}

// NewMyEventReceiver 创建一个新的 MyEventReceiver
func NewMyEventReceiver(logger *log.Logger) *MyEventReceiver {
	return &MyEventReceiver{logger: logger}
}

// Event 记录基本的事件信息
func (r *MyEventReceiver) Event(eventName string) {
	r.logger.Printf("Event: %s\n", eventName)
}

// EventKv 记录带键值对的事件信息
func (r *MyEventReceiver) EventKv(eventName string, kvs map[string]string) {
	r.logger.Printf("Event: %s, KVs: %v\n", eventName, kvs)
}

// EventErr 记录事件和错误信息
func (r *MyEventReceiver) EventErr(eventName string, err error) error {
	r.logger.Printf("Event: %s, Error: %v\n", eventName, err)
	return err
}

// EventErrKv 记录事件、错误以及键值对
func (r *MyEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	r.logger.Printf("Event: %s, Error: %v, KVs: %v\n", eventName, err, kvs)
	return err
}

// Timing 记录事件的耗时（纳秒）
func (r *MyEventReceiver) Timing(eventName string, nanoseconds int64) {
	duration := time.Duration(nanoseconds) * time.Nanosecond
	r.logger.Printf("Event: %s took %v\n", eventName, duration)
}

// TimingKv 记录事件的耗时并附带键值对
func (r *MyEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	duration := time.Duration(nanoseconds) * time.Nanosecond
	r.logger.Printf("Event: %s took %v, KVs: %v\n", eventName, duration, kvs)
}

// 示例：使用 MyEventReceiver 实现
func main() {
	// 创建一个 logger 用于记录日志
	logger := log.New(log.Writer(), "EventReceiver: ", log.LstdFlags)

	// 初始化 MyEventReceiver
	receiver := NewMyEventReceiver(logger)

	// 调用不同的事件方法
	receiver.Event("user_login")
	receiver.EventKv("order_created", map[string]string{"order_id": "12345", "user_id": "67890"})
	receiver.EventErr("payment_failed", fmt.Errorf("Insufficient funds"))
	receiver.EventErrKv("payment_failed", fmt.Errorf("Insufficient funds"), map[string]string{"order_id": "12345", "user_id": "67890"})
	receiver.Timing("db_query", 1500000000)
	receiver.TimingKv("db_query", 1500000000, map[string]string{"query": "SELECT * FROM users"})
}
