package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	{
		ctx1, cancel1 := context.WithCancel(context.Background())
		ctx2, cancel2 := context.WithCancel(ctx1)
		ctx3, cancel3 := context.WithCancel(ctx2)
		ctx4, cancel4 := context.WithCancel(ctx3)
		defer cancel1()
		defer cancel3()
		defer cancel4()

		cancel2()
		//当cancel2()执行后，此ctx2，以及ctx2的后代ctx3,ctx4都会被关闭。
		//但是ctx2的父代ctx1没事
		fmt.Println("After cancel2()")
		fmt.Println("ctx1.Err()=", ctx1.Err())
		fmt.Println("ctx2.Err()=", ctx2.Err())
		fmt.Println("ctx3.Err()=", ctx3.Err())
		fmt.Println("ctx4.Err()=", ctx4.Err())

		fmt.Println("=================================")
	}
	{
		//当子ctx的daealine在父ctx之后时，子ctx沿用父ctx的deadline

		timeout1_6 := 6 * time.Second
		timeout2_4 := 4 * time.Second
		timeout3_8 := 8 * time.Second

		ctx1, cancel1 := context.WithTimeout(context.Background(), timeout1_6)
		ctx2, cancel2 := context.WithTimeout(ctx1, timeout2_4)
		ctx3, cancel3 := context.WithTimeout(ctx2, timeout3_8)
		defer cancel1()
		defer cancel2()
		defer cancel3()

		deadline1, ok1 := ctx1.Deadline()
		if ok1 {
			fmt.Println("ctx1.Deadline()=", deadline1)
		}

		deadline2, ok2 := ctx2.Deadline()
		if ok2 {
			fmt.Println("ctx2.Deadline()=", deadline2)
			fmt.Println("可以看到，若ctx的deadline早于父ctx，可以设置成功。")
		}

		deadline3, ok3 := ctx3.Deadline()
		if ok3 {
			fmt.Println("ctx3.Deadline()=", deadline3)
			fmt.Println("可以看到，若ctx的deadline 晚于 父ctx，则沿用父ctx的deadline。")
		}

		fmt.Println("=================================")
	}

	{
		//子ctx    能访问父ctx的值
		//父ctx 不 能访问子ctx的值

		ctx1 := context.WithValue(context.Background(), 1, 1)
		ctx2 := context.WithValue(ctx1, 2, 2)

		value1_2 := ctx1.Value(2)
		fmt.Println("父ctx无法访问子ctx的值，ctx1访问ctx2的赋值=", value1_2)

		value2_1 := ctx2.Value(1)
		fmt.Println("子ctx可以访问父ctx的值，ctx2访问ctx1的赋值=", value2_1)

		fmt.Println("=================================")
	}
}
