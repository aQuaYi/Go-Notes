package main

import (
	"context"
	"log"
	"time"
)

func main() {
	cancel := func(funcName string, cf context.Context) {
		<-cf.Done()
		log.Printf("%s Canceled: %s", funcName, cf.Err())
	}

	c1, cf1 := context.WithCancel(context.Background())
	go cancel("c1", c1)

	c2, cf2 := context.WithCancel(c1)
	go cancel("c2", c2)

	c3, cf3 := context.WithCancel(c2)
	go cancel("c3", c3)

	// 取消下一行的注释，观察输出顺序的变化
	// cf1()

	time.Sleep(time.Millisecond * 50)
	cf3()

	time.Sleep(time.Millisecond * 50)
	cf2()

	time.Sleep(time.Millisecond * 50)
	cf1()

	time.Sleep(time.Millisecond * 50)
}
