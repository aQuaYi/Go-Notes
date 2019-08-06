package orderprint

import (
	"runtime"
	"sync"
)

type chanWorker struct {
	count int
}

func newChanWorker() *chanWorker {
	w := &chanWorker{}
	return w
}

// Work 触发事件
func (w *chanWorker) Work() {
	w.count++
}

func channel() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	var wg sync.WaitGroup

	w1 := newChanWorker()
	go fakeChannel(&wg, w1.Work, c1, c2)

	w2 := newChanWorker()
	go fakeChannel(&wg, w2.Work, c2, c3)

	w3 := newChanWorker()
	go fakeChannel(&wg, w3.Work, c3, c1)

	c1 <- 300
	wg.Wait()

	// fmt.Println(w1.count, w2.count, w3.count)
}

func fakeChannel(wg *sync.WaitGroup, work func(), cur, next chan int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	(*wg).Add(1)
	defer (*wg).Done()
	for {
		i := <-cur
		if i == 0 {
			close(next)
			return
		}
		work()
		next <- i - 1
	}
}
