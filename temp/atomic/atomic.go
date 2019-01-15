package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Hello World")

	var wg sync.WaitGroup
	wg.Add(1)

	num2 := int32(1)

	go func() {
		for {
			if atomic.CompareAndSwapInt32(&num2, 10, 0) {
				fmt.Println("The second number has gone to zero.")
				break
			}
			time.Sleep(time.Millisecond * 50)
		}
		wg.Done()
	}()

	time.Sleep(time.Second)
	num2 = 10
	wg.Wait()
	fmt.Println(num2)
}
