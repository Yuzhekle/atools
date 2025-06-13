package main

import (
	"io"
	"net"
	"os"
)

func main() {
	file, _ := os.Open("source.txt")
	defer file.Close()
	conn, _ := net.Dial("tcp", "example.com:80")
	io.Copy(conn, file) // Go 的 io.Copy 内部使用 sendfile（若支持）
}
