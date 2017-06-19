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
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		fmt.Println("ctx.Err()     =", ctx.Err())
		fmt.Println("time.Now()    =", time.Now())
		dl, ok := ctx.Deadline()
		if ok {
			fmt.Println("ctx.Deadline()=", dl)
		}
		res := add(ctx, a, b)
		fmt.Printf("%d+%d=%d\n", a, b, res)
		fmt.Println(ctx.Err())
	}

	{
		//使用cancel终结
		ctx, cancel := context.WithCancel(context.TODO())
		go func() {
			time.Sleep(2 * time.Second)
			cancel()
		}()

		res := add(ctx, a, b)
		fmt.Printf("%d+%d=%d\n", a, b, res)
		fmt.Println(ctx.Err())
	}

}
