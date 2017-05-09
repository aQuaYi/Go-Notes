package primeNumber

import "sync"
import "time"

var wg sync.WaitGroup

//发送所有的整数
func generate(ch chan<- int) {
	time.Sleep(time.Millisecond * 100)
	for i := 2; i < 1000; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { //src发送来的，是经过前面所有素数筛选后的数。
		if i%prime != 0 {
			dst <- i //成功通过此prime检验的i会被发送给下一个prime检验
			//特别地，第一个通过此检验的i，本身就是素数。
		}
	}
	close(dst)
	wg.Done()
}

func answer(src <-chan int) <-chan int {
	result := make(chan int)

	go func() {

		for {
			wg.Add(1)
			dst := make(chan int)
			prime := <-src
			result <- prime
			go filter(src, dst, prime)
			src = dst
		}
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

//Chan 返回一个发送prime number的channel
func Chan() <-chan int {
	src := make(chan int) // Create a new channel.
	wg.Add(1)
	go generate(src) // Start generate() as a subprocess.
	return answer(src)
}
