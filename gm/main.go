package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var m = map[string]int{"a": 1, "b": 2, "c": 3}

func main() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			t := time.NewTicker(5 * time.Second)
			s := run(m, i)
			fmt.Println("s", s)
			println()

			for {
				select {
				case <-t.C:
					clear()
					println("===========================================")
					s = run(m, i)
					fmt.Println("s", s)
					println()
					println()
				}
			}
		}(i)
	}

	select {}
}

func run(m map[string]int, i int) []int {
	println("i", i)
	s := make([]int, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}

	return s
}

func clear() {
	// 声明一个Command对象，用于执行命令
	cmd := exec.Command("clear")

	// 将命令标准输出连接到当前进程的标准输出
	cmd.Stdout = os.Stdout

	// 执行命令
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
