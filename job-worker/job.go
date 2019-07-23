package main

import (
	"fmt"
	"math"
)

// intStream: a buffered channel that we can send work requests on.
var intStream chan int

// Worker represents the worker that executes the job
type Worker struct {
	ID      int
	pool    chan chan int
	intChan chan int
	quit    chan struct{}
}

func newWorker(id int, workerPool chan chan int) Worker {
	return Worker{
		ID:      id,
		pool:    workerPool,
		intChan: make(chan int),
		quit:    make(chan struct{})}
}

func (w Worker) start() {
	go func() {
		for {
			// 已空闲，把获取通道送还到工作池
			w.pool <- w.intChan
			select {
			case num := <-w.intChan:
				fmt.Printf("[ID:%d] receive %d, ", w.ID, num)
				if isPrime(num) {
					fmt.Println("It's prime")
				} else {
					fmt.Println("It's NOT prime")
				}
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		close(w.quit)
	}()
}

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n)))
	for i := 0; i <= root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
