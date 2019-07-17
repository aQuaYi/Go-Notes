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
	nums ...int,
) (
	<-chan interface{},
	<-chan int,
) {
	heartbeat := make(chan interface{}, 1) // TODO: 取消缓存看看效果
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)
		// 模拟各种延迟，比如 网络延迟，goroutine执行延迟等
		x := time.Duration(rand.Intn(2000))
		time.Sleep(x * time.Millisecond)
		//
		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
			select {
			case <-done:
				return
			case intStream <- n:
			}
		}
	}()

	return heartbeat, intStream
}
