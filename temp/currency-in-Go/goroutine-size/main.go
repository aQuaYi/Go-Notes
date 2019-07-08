package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}

	var wg sync.WaitGroup

	noExit := func() {
		wg.Done()
		<-c // no data out
	}

	const numGoroutines = 1e8
	wg.Add(numGoroutines)

	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noExit()
	}
	wg.Wait()
	after := memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
