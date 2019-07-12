package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	unstableChan := makeSleepChan()
	for v := range orDone(
		time.After(time.Second*12),
		unstableChan,
	) {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// done 给 c 添加了抢占功能
// 避免了阻塞
var orDone = func(done <-chan time.Time, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case valStream <- v:
				}
			}
		}
	}()
	return valStream
}

func makeSleepChan() <-chan interface{} {
	res := make(chan interface{})
	go func() {
		defer close(res)
		for i := 0; i < 12; i++ {
			res <- i
			randDuration := time.Duration(rand.Intn(3)) * time.Second
			time.Sleep(randDuration)
		}
		fmt.Println("All send")
	}()
	return res
}
