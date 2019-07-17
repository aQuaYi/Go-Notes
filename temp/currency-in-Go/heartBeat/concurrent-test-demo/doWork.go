package work

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func doWork(
	done <-chan interface{},
	pulseInterval time.Duration,
	nums ...int,
) (
	<-chan interface{},
	<-chan int,
) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)
		// 模拟各种延迟，比如 网络延迟，goroutine执行延迟等
		x := time.Duration(rand.Intn(2000))
		time.Sleep(x * time.Millisecond)
		//
		pulse := time.Tick(pulseInterval)
	numLoop:
		for _, n := range nums {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case intStream <- n:
					continue numLoop
				}
			}
		}
	}()
	//
	return heartbeat, intStream
}
