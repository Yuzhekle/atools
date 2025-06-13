package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wrong3()
}

func wrong3() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 原始 Pool
	var p1 = sync.Pool{
		New: func() any {
			fmt.Println("p1.New: allocating new object")
			return make([]byte, 1024)
		},
	}

	// 向 p1 预填 1000 个对象
	for i := 0; i < 1000; i++ {
		p1.Put(make([]byte, 1024))
	}

	// ❌ 复制了 Pool：p2 拥有自己的缓存区
	p2 := p1

	var wg sync.WaitGroup

	// 使用 p2 的 goroutines（错误地以为能复用 p1 的对象）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				obj := p2.Get().([]byte)
				time.Sleep(1 * time.Millisecond)
				p2.Put(obj)
			}
		}(i)
	}

	wg.Wait()
}

func wrong2() {
	var p1 = sync.Pool{
		New: func() any {
			fmt.Println("p1.New: allocating new object")
			return "new-object"
		},
	}

	// 正确使用：直接用 p1.Put / p1.Get
	p1.Put("hello-from-p1")

	// ❌ 错误使用：复制 p1 到 p2
	p2 := p1

	// 观察 p1 的 Get：能命中缓存
	fmt.Println("p1.Get():", p1.Get()) // 输出 hello-from-p1

	// 观察 p2 的 Get：缓存是空的，触发 New
	fmt.Println("p2.Get():", p2.Get()) // 输出 new-object（调用了 New）

	// 验证是否是两个 pool 的缓存不同（Put 到 p2，再从 p1 Get）
	p2.Put("written-to-p2")

	// 再次从 p1 Get，是否拿得到？一般情况是拿不到
	fmt.Println("p1.Get():", p1.Get()) // 很可能还是触发 New
}

func wrong1() {
	// p1 := sync.Pool{}
	// p2 := p1

	// p2.Put("Hello")

	// fmt.Println(p2.Get())

	// p3 := P{}
	// p4 := p3

	// fmt.Println(p4)

	var p1 sync.Pool
	p1.New = func() any {
		return make([]byte, 1024)
	}

	p1.Put([]byte("abc"))

	p2 := p1

	go func() {
		for i := 0; i < 1000; i++ {
			p2.Put([]byte("abc"))
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(p1.Get())
		}
	}()
}

type P struct {
	// noCopy noCopy
}

type noCopy struct{}

func (*noCopy) Lock() {}

func (*noCopy) Unlock() {}
