package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Println(id, ": ready.")
			<-ready
			fmt.Println(id, ": running...")
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Ready?...Go...")

	close(ready)

	wg.Wait()
}
