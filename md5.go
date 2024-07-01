package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	fmt.Printf("%b,%b, %b\n\n", uint8(1), -1, ^int8(1))
	fmt.Printf("%b,%b, %b, %b\n\n", 64, 63, ^63, 64&^63)
	fmt.Println(64 &^ 63)

	msg := "Hello, World!你好世界"
	for i := 0; i < 10; i++ {
		h := md5.New()
		h.Write([]byte(msg))
		s := hex.EncodeToString(h.Sum(nil))
		println(i, string(s))
	}
}
