package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	example()
}

func example() {
	go func() {
		if err := recover(); err != nil {
			fmt.Println(1)
		}
	}()
	go a()
	go func() {
		if err := recover(); err != nil {
			fmt.Println(2)
		}
	}()
}

func a() {
	fmt.Println(1 / 0)
}

func tt() {
	// 异步获取订单的预测取消概率
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resultChan := make(chan float64, 1) // 有缓冲channel防止goroutine泄漏

	go func() {
		probability, err := testFunc()
		if err != nil {
			probability = 0 // 使用默认值
		}

		resultChan <- probability
		// select {
		// case resultChan <- probability:
		// 	println(111)
		// case <-ctx.Done(): // 超时不再发送
		// 	resultChan <- 0
		// 	println(222)
		// }
	}()

	// probability := <-resultChan

	var probability float64
	select {
	case probability = <-resultChan:
		// cancel()
		println(444)
	case <-ctx.Done():
		probability = 0 // 超时后使用默认值
		println(333)
	}

	// println(555)

	println(probability)

}

func testFunc() (float64, error) {
	// panic("test panic")
	time.Sleep(1 * time.Second)
	return 0, nil
}
