package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("Ward: Hello, I'm irresponsible!")
		go func() {
			<-done
			log.Println("Ward: I am halting.")
		}()
		return nil
	}

	doWorkWithSteward := newSteward(4*time.Second, doWork)

	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for range doWorkWithSteward(done, 4*time.Second) {
	}
	log.Println("Done.")
}

type startGoroutineFn func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (heartbeat <-chan interface{})

var newSteward = func(
	timeout time.Duration,
	startGoroutine startGoroutineFn,
) startGoroutineFn {
	return func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) <-chan interface{} { // 2
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}
			startWard := func() { // 3
				wardDone = make(chan interface{})                             // 4
				wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2) //5
			}
			startWard()
			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)

				for { // 6
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()
		return heartbeat
	}
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
