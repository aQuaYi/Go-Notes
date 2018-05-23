package printIn2Goroutine

import (
	"runtime"
	"sync"
)

func waitGroupPrinter() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var num, char, wg sync.WaitGroup
	num.Add(1)
	wg.Add(2)

	go func() {
		for i := 1; i <= 10; i += 2 {
			char.Wait()
			char.Add(1)
			for j := i; j <= 10 && j < i+2; j++ {
				//fmt.Print(j)
			}
			num.Done()
		}
		wg.Done()
	}()

	go func() {
		s := "ABCDEFGHIJ"
		for i := 0; i < len(s); i += 2 {
			num.Wait()
			num.Add(1)
			for j := i; j < len(s) && j < i+2; j++ {
				//fmt.Print(s[j : j+1])
			}
			char.Done()
		}
		wg.Done()
	}()

	wg.Wait()
}

func channelPrinter() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	chan_n := make(chan struct{})
	chan_c := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for i := 1; i < 11; i += 2 {
			<-chan_c
			//fmt.Print(i)
			//fmt.Print(i + 1)
			if i+2 >= 11 {
				close(chan_n)
				return
			}
			chan_n <- struct{}{}
		}
	}()

	go func() {
		//char_seq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-chan_n
			//fmt.Print(char_seq[i])
			//fmt.Print(char_seq[i+1])
			if i+2 >= 10 {
				close(chan_c)
				close(done)
				return
			}
			chan_c <- struct{}{}
		}
	}()

	chan_c <- struct{}{}
	<-done
}
