package main

import (
	"strconv"
)

func main() {
	s := strconv.FormatInt(4095, 16)
	println(s)
	s1 := strconv.FormatInt(65535, 20)
	println(s1)
	println(4096 * 16)
}
