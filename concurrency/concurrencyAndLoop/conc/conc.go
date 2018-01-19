package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	n := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS = ", n)
	conc(n)
}

func count() {
	x := 0
	for i := 0; i < 1<<34-1; i++ {
		x += i
	}
	fmt.Println(x)
}

func conc(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {

		go func() {
			count()
			wg.Done()
		}()
	}

	wg.Wait()
}
