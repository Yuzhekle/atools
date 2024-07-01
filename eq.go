package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [3]int{5, 6}
	b := [2]int{5, 6}
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	r := "hello"
}
