package prime

import (
	"math"
	"sync"
)

type prime struct {
	ch    chan int
	limit int
	wg    sync.WaitGroup
}

//NewUnder 设置了产生素数的范围
func NewUnder(limit int) <-chan int {
	c := make(chan int)
	p := &prime{
		ch:    c,
		limit: limit,
	}

	src := p.generate()
	p.loop(src)

	return p.ch
}

//New 会发送int范围内的所有素数
func New() <-chan int {
	return NewUnder(math.MaxInt64)
}

func (p *prime) generate() <-chan int {
	src := make(chan int)
	p.wg.Add(1)

	go func() {
		//发送所有在limit范围内的数。
		for i := 2; i <= p.limit; i++ {
			src <- i
		}
		//发送完毕后，关闭通道
		close(src)
		p.wg.Done()
	}()

	return src
}

func (p *prime) filter(src <-chan int, prime int) <-chan int {
	dst := make(chan int)
	p.wg.Add(1)

	go func() {
		//利用此prime对已经前面筛选步骤的数进行进一步的筛选。
		for i := range src { //src发送来的，是通过了前面所有素数筛选的数。
			if i%prime != 0 {
				dst <- i //成功通过此prime检验的i会被发送给下一个prime检验
				//特别地，第一个通过此检验的i，是下一个prime
			}
		}
		//发送完成后，关闭发送通道
		close(dst)
		p.wg.Done()
	}()

	return dst
}

func (p *prime) loop(src <-chan int) {
	go func() {
		for {
			prime, isOpen := <-src //每个src发出的第一个数，都是素数。
			if !isOpen {
				//最后一个src(下文称为src-1)，直到倒数第二个src(下文称为src-2)关闭，都没有来得及发现下一个素数。但是由于src-2关闭了，src-1也马上被关闭了。此时，反而可以从src-1中取值了。只是，此时的isOpen被为false。经过if语句判断后，会结束本for循环。
				break
			}
			p.ch <- prime
			src = p.filter(src, prime)
		}
		//等待所有的src关闭
		p.wg.Wait()
		close(p.ch)
	}()
}
