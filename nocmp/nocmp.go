package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func main() {
	t := &T{}
	t.Set("hello")
	println(t.Get())

	t1 := &T{}
	t1.Set("world")
	println(t1.Get())

	println(t == t1)

	var s = []int{1, 5, 4, 3, 2, 9, 6, 7}
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	fmt.Println("s", s)
}

// nocmp is a type that does not implement the comparison operators.
type nocmp [0]func()

type T struct {
	_ nocmp

	v string
}

func (t *T) Set(v string) {
	t.v = v
}

func (t *T) Get() string {
	return t.v
}
