package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	rsCh := make(chan int, 1)
	stopCh := make(chan struct{})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rsCh <- i
		}(i)
	}

	go func() {
		for {
			select {
			case rs := <-rsCh:
				println(rs)
			case <-stopCh:
				return
			}
		}
	}()

	wg.Wait()
	// close(rsCh)
	// time.Sleep(1 * time.Second)
	close(stopCh)

	// select {}
}
