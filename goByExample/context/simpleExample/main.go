package main

import (
	"context"
	"fmt"
	"time"
)

func inc(a int) int {
	res := a + 1
	time.Sleep(1 * time.Second)
	return res
}

func add(ctx context.Context, a, b int) int {
	res := 0

	for i := 0; i < a; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}

	for i := 0; i < b; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}

	return res
}

func main() {
	a, b := 3, 2

	{
		//使用Timeout终结
		timeout := 2 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		fmt.Println(ctx.Err())
		res := add(ctx, a, b)
		fmt.Printf("%d+%d=%d\n", a, b, res)
		fmt.Println(ctx.Err())
		defer cancel()
	}

	{ // 使用大括号，感觉很有层次
		//使用cancel终结
		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			time.Sleep(2 * time.Second)
			cancel()
		}()

		res := add(ctx, a, b)
		fmt.Printf("%d+%d=%d\n", a, b, res)
		fmt.Println(ctx.Err())
	}
}
