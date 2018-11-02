package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := range ch {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("ch已经被关闭。")
		wg.Done()
	}()

	wg.Wait()
}
