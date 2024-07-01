package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	GetTaoBaoAPISetting()
}

var onceQuery = sync.Once{}

func GetTaoBaoAPISetting() {
	onceQuery.Do(func() {
		go func() {
			ticker := time.NewTicker(time.Second * 3)
			defer ticker.Stop()

			for range ticker.C {
				fmt.Println("Hello, 世界")
			}
		}()

		fmt.Println("Hello, once")
	})
}
