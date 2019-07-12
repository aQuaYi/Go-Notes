package main

import (
	"fmt"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}), 7)
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 2)
				stream <- i
				close(stream)
				chanStream <- stream
			}
			fmt.Println("\n All send.")
		}()
		return chanStream
	}
	for v := range bridge(nil, genVals()) {
		fmt.Printf(" %v ", v)
	}
	fmt.Println()
}

var bridge = func(
	done <-chan interface{},
	chanStream <-chan <-chan interface{},
) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		var c <-chan interface{}
		var ok bool
		for {
			select {
			case <-done:
				return // 严防死守
			case c, ok = <-chanStream:
				if !ok {
					return // 严防死守
				}
			}
			for v := range orDone(done, c) {
				select {
				case <-done:
					return // 严防死守
				case valStream <- v:
				}
			}
		}
	}()
	return valStream
}

var orDone = func(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return // 严防死守
			case v, ok := <-c:
				if !ok {
					return // 严防死守
				}
				select {
				case <-done:
					return // 严防死守
				case valStream <- v:
				}
			}
		}
	}()
	return valStream
}
