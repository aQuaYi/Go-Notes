package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	<-or(
		sig(10*time.Second),
		sig(1*time.Second),
		sig(1*time.Minute),
		sig(2*time.Hour),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

// done 用于关闭 or goroutine
// channels 中的任意一个接收到信号后，都会通过 orDone 的关闭来转发
func or(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return done
	}
	//
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		select {
		case <-done:
		case <-channels[0]:
		case <-or(orDone, channels[1:]...):
		}
	}()
	//
	return orDone
}
