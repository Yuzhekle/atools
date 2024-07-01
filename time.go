package main

import (
	"fmt"
	"time"
)

func init() {
	run()
}

func main() {

	for {
	}
}

func run() {
	go func() {
		for {
			Update()
			time.Sleep(1 * time.Second)
		}
	}()
}

func Update() {
	fmt.Println("update")
}
