package primeNumber

import "math"

//发送所有的整数
func generate(ch chan<- int) {
	for i := 2; i < math.MaxInt64; i++ {
		ch <- i
	}
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { //src发送来的，是经过前面所有素数筛选后的数。
		if i%prime != 0 {
			dst <- i //成功通过此prime检验的i会被发送给下一个prime检验
			//特别地，第一个通过此检验的i，本身就是素数。
		}
	}
}

func answer(src <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for {
			prime := <-src
			result <- prime
			dst := make(chan int)
			go filter(src, dst, prime)
			src = dst
		}
	}()
	return result
}

//Chan 返回一个发送prime number的channel
func Chan() <-chan int {
	src := make(chan int) // Create a new channel.
	go generate(src)      // Start generate() as a subprocess.
	return answer(src)
}
