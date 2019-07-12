package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	var wg sync.WaitGroup

	output := func(name string, c <-chan interface{}) {
		defer wg.Done()
		for v := range c {
			fmt.Printf("%s receive %v\n", name, v)
		}
	}

	out1, out2 := tee(done, makeIntChan())

	wg.Add(2)

	go output("1", out1)
	go output("2", out2)

	wg.Wait()

	fmt.Println("Done")
}

// 将 in 中的流分叉成两条 out
// 两个 out 中的数值是一样的
var tee = func(
	done <-chan interface{},
	in <-chan interface{},
) (_, _ <-chan interface{}) { // <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range orDone(done, in) {
			o1, o2 := out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case o1 <- val:
					o1 = nil
				case o2 <- val:
					o2 = nil
				}
			}
		}
	}()
	return out1, out2
}

var orDone = func(done, c <-chan interface{}) <-chan interface{} {
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

func makeIntChan() <-chan interface{} {
	intStream := make(chan interface{})
	go func() {
		defer close(intStream)
		for i := 0; i < 12; i++ {
			intStream <- i
		}
	}()
	return intStream
}
