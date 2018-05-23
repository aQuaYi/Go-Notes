package prime

import "math"

//NewUnder 返回一个channel 发送 all prime < limit
func NewUnder(limit int) <-chan int {
	src := generate(limit)
	result := loop(src)
	return result
}

//New 需要设置产生素数的范围
func New() <-chan int {
	return NewUnder(math.MaxInt64)
}

func generate(limit int) <-chan int {
	src := make(chan int)

	go func() {
		//发送所有在limit范围内的数。
		for i := 2; i < limit; i++ {
			src <- i
		}
		//发送完毕后，关闭通道
		close(src)
	}()

	return src
}

func filter(src <-chan int, prime int) <-chan int {
	dst := make(chan int)

	go func() {
		//利用此prime对已经前面筛选步骤的数进行进一步的筛选。
		for i := range src { //src发送来的，是通过了前面所有素数筛选的数。
			if i%prime != 0 {
				dst <- i //成功通过此prime检验的i会被发送给下一个prime检验
				//特别地，第一个通过此检验的i，是下一个prime
			}
		}
		//发送完成后，关闭通道
		//结束下一个filter中的for循环
		close(dst)
	}()

	return dst
}

func loop(src <-chan int) <-chan int {
	result := make(chan int)

	go func() {

		for {
			prime, isOpen := <-src //每个src发出的第一个数，都是素数。
			if !isOpen {
				//最后一个src(下文称为src-1)，直到倒数第二个src(下文称为src-2)关闭，都没有来得及发现下一个素数。但是由于src-2关闭了，src-1也马上被关闭了。此时，反而可以从src-1中取值了。只是，此时的isOpen为false。经过if语句判断后，会结束本for循环。
				break
			}
			result <- prime
			src = filter(src, prime)
		}

		close(result)
	}()

	return result
}
