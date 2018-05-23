package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) // 最多允许 2 个任务并发

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        // 获取型号
			defer func() { <-sem }() // 释放信号

			fmt.Println(id, time.Now())
			time.Sleep(time.Second * 2)
		}(i)
	}

	wg.Wait()
}
