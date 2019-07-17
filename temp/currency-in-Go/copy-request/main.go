package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	done := make(chan interface{})

	workers := 10
	var wg sync.WaitGroup
	wg.Add(workers)

	result := make(chan int)
	for i := 0; i < workers; i++ {
		go doWork(done, i, &wg, result)
	}

	finished := <-result

	close(done)
	wg.Wait()

	fmt.Printf("finished by %d\n", finished)
}

func doWork(
	done <-chan interface{},
	id int,
	wg *sync.WaitGroup,
	result chan<- int,
) {
	started := time.Now()
	defer wg.Done()

	// 模拟随机负载
	simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
	finished := "finished by me"
	select {
	case <-done:
		finished = "finished by the other"
	case <-time.After(simulatedLoadTime):
	}
	select {
	case <-done:
		finished = "finished by the other"
	case result <- id:
	}

	fmt.Printf("%d need %v, but took %v, %s\n",
		id,
		simulatedLoadTime,
		time.Since(started),
		finished)
}
