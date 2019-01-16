package main

import (
	"context"
	"fmt"
)

func main() {
	for {
		// 运行的时候，可以打开资源管理器，观察内存使用率
		goroutineLeak()
		// withoutLeak()
	}
}

func goroutineLeak() {
	// 由于 gen 生成的 goroutine 都没有关闭
	// 所以会造成 goroutine leak
	gen := func() <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				dst <- n
				n++
			}
		}()
		return dst
	}

	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func withoutLeak() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
