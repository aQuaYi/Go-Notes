package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		log.Println("close done")
		close(done)
	})

	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			log.Println("Pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			log.Printf("result: %v\n", r.Second())
		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}
}

var doWork = func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (
	<-chan interface{},
	<-chan time.Time,
) {
	heartbeat := make(chan interface{})
	results := make(chan time.Time)

	go func() {
		defer close(heartbeat)
		defer close(results)

		pulse := time.Tick(pulseInterval)
		workGen := time.Tick(2 * pulseInterval)

		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}

		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workGen:
				sendResult(r)
			}
		}
	}()

	return heartbeat, results
}
