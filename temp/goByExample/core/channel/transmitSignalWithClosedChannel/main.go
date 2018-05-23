package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	finish := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		waiter(finish, i, wg)
	}

	go func() {
		// channel 关闭后，仍然可以从channel中获取通道数据类型的零值
		// 同时，获取的验证数据为false，表示获取的不是真正从通道中传递过来的值
		f, ok := <-finish
		fmt.Println("f  = ", f)
		fmt.Println("ok = ", ok)
		// channel 关闭后，任然能从channel中获取到零值。
		// 利用这样的方式，
		// 不是在利用 channel 传递data，
		// 而是在利用 channel的关闭 传递signal
		// 也就是说，传递的内容不重要，传递过来一个东西，让我知道有事发生，就好了。
		fmt.Println("可以看到，所有的waiter几乎同时接收到了finish信号。")
		// 这种方式最重要的功能是，不会由于某个获取阻塞了，而影响到其他地方的获取。
	}()

	time.Sleep(time.Second)
	close(finish)

	wg.Wait()
}

func waiter(finish chan int, i int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		<-finish
		log.Printf("%02d waiter, finished", i)
		wg.Done()
	}()
}
