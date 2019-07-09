package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		front := queue[0]
		queue = queue[1:]
		fmt.Println("\t\tRemoved from queue", front)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 100; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue", i)
		queue = append(queue, i)
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
