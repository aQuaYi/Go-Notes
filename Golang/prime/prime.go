package prime

import (
	"math"
	"sync"
)

//Prime 是以发送素数的channel为主的结构体
type Prime struct {
	Chan  chan int
	limit int
	wg    sync.WaitGroup
}

//NewUnder 设置了产生素数的范围
func NewUnder(limit int) <-chan int {
	c := make(chan int)
	p := &Prime{
		Chan:  c,
		limit: limit,
	}

	src := p.generate()
	p.loop(src)

	return p.Chan
}

//New 会发送int范围内的所有素数
func New() <-chan int {
	return NewUnder(math.MaxInt64)
}

func (p *Prime) generate() <-chan int {
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

func (p *Prime) filter(src <-chan int, prime int) <-chan int {
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

func (p *Prime) loop(src <-chan int) {
	go func() {

		for {
			prime, isOpen := <-src
			if !isOpen {
				break
			}
			p.Chan <- prime
			src = p.filter(src, prime)
		}
		//等待所有的src关闭
		p.wg.Wait()
		close(p.Chan)
	}()
}
