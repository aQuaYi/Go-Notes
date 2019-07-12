package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	randInt := func() interface{} { return rand.Intn(500000000) + 500000000 }
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	randintStream := toInt(done, repeatFn(done, randInt))
	fan := fanIn(done, fanOut(done, randintStream, primeFinder))
	pipeline := take(done, fan, 10)
	fmt.Println("Primes:")
	for prime := range pipeline {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v\n", time.Since(start))
}

var fanOut = func(
	done <-chan interface{},
	intStream <-chan int,
	fanStage func(<-chan interface{}, <-chan int) <-chan interface{},
) []<-chan interface{} {
	num := runtime.NumCPU()
	finders := make([]<-chan interface{}, num)
	for i := 0; i < num; i++ {
		finders[i] = fanStage(done, intStream)
	}
	return finders
}

var fanIn = func(
	done <-chan interface{},
	channels []<-chan interface{},
) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

var repeatFn = func(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

var toInt = func(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

var primeFinder = func(
	done <-chan interface{},
	intStream <-chan int,
) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for i := range intStream {
			if !isPrime(i) {
				continue
			}
			select {
			case <-done:
				return
			case primeStream <- i:
			}
		}
	}()
	return primeStream
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var take = func(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}
