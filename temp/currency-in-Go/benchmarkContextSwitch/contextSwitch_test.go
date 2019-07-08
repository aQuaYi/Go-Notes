package benchmarkContextSwitch

import (
	"sync"
	"testing"
)

func Benchmark_contextSwitch(b *testing.B) {
	var wg sync.WaitGroup

	begin := make(chan struct{})
	c := make(chan struct{})

	sender := func() {
		<-begin
		for i := 0; i < b.N; i++ {
			c <- struct{}{}
		}
		wg.Done()
	}

	receiver := func() {
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
		wg.Done()
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
