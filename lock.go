package main

import "sync"

func main() {
	s := sync.Mutex{}
	s.Lock()
	defer s.Unlock()
}
