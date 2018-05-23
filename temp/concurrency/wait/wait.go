package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	begin := time.Now()

	go func() {
		wg.Wait() // 所有等待的地方，都会收到通知
		fmt.Println("Wait in go func", time.Since(begin))
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second * time.Duration(rand.Intn(id+1)))
			fmt.Println("ID:", id, "Done.", time.Since(begin))
		}(i)
	}

	fmt.Println("main......", begin)
	wg.Wait()
	fmt.Println("main exit.", time.Since(begin))
}
