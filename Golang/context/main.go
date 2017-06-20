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
		timeout2 := 2 * time.Second
		timeout5 := 5 * time.Second
		ctx2, cancel2 := context.WithTimeout(context.Background(), timeout2)
		ctx5, cancel5 := context.WithTimeout(ctx2, timeout5)
		fmt.Println("ctx.Err()     =", ctx2.Err())
		fmt.Println("time.Now()    =", time.Now())
		dl2, ok := ctx2.Deadline()
		if ok {
			fmt.Println("ctx2.Deadline()=", dl2)
		}
		dl5, ok := ctx5.Deadline()
		if ok {
			fmt.Println("ctx5.Deadline()=", dl5)
		}
		res := add(ctx2, a, b)
		fmt.Printf("%d+%d=%d\n", a, b, res)
		fmt.Println(ctx2.Err())
		cancel2()
		fmt.Println("After cancel2(), ctx5.Err()=", ctx5.Err())
		cancel5()

	}

	{
		//使用cancel终结
		timeout5 := 5 * time.Second
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		ctx5, cancel5 := context.WithTimeout(ctx, timeout5)

		fmt.Println("time.Now()    =", time.Now())
		dl5, ok := ctx5.Deadline()
		if ok {
			fmt.Println("ctx5.Deadline()=", dl5)
		}

		go func() {
			time.Sleep(2 * time.Second)
			cancel5()
		}()

		res := add(ctx, a, b)
		fmt.Printf("%d+%d=%d\n", a, b, res)
		fmt.Println(ctx.Err())
	}
	{
		//当ctx对应的cancel()执行后，此ctx及其所有子ctx都会被关闭。
		//但是此ctx的父ctx没事。

		ctx1, cancel1 := context.WithCancel(context.Background())
		ctx2, cancel2 := context.WithCancel(ctx1)
		ctx3, cancel3 := context.WithCancel(ctx2)
		ctx4, cancel4 := context.WithCancel(ctx3)
		defer cancel1()
		defer cancel2()
		defer cancel4()

		cancel3()
		fmt.Println("After cancel3()")
		fmt.Println("ctx1.Err()=", ctx1.Err())
		fmt.Println("ctx2.Err()=", ctx2.Err())
		fmt.Println("ctx3.Err()=", ctx3.Err())
		fmt.Println("ctx4.Err()=", ctx4.Err())
	}

}
