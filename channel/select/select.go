package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	a := make(chan int)
	b := make(chan int)

	go func() { // 发送端
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)

			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}

			if !ok {
				return
			}

			fmt.Println(name, ":", x)

			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()

	wg.Wait()
}
