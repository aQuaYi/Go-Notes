package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	for i := 0; i < 100; i++ {
		increments.Add(1)
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Println(count)
}
