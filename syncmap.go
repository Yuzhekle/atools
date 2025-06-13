package main

import (
	"fmt"
	"sync"
)

func main() {

	smap := sync.Map{}
	smap.Store("key", "value")

	smap.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	m := map[string]string{}
	m["key"] = "value"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
