package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// mailbox 代表信箱。
	var mailbox string

	// lock 代表信箱上的锁。
	var lock sync.Mutex

	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)

	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(&lock)

	max := 5
	workers := 5

	// 用于发信。
	for w := 0; w < workers; w++ {
		go func(id int) {
			for i := 1; i <= max; i++ {
				sendCond.L.Lock()
				for mailbox != "" { // [1]
					sendCond.Wait()
				}
				mailbox = fmt.Sprintf("S%d-%d", id, i)
				sendCond.L.Unlock()
				recvCond.Signal()
			}
		}(w)
	}

	// 用于收信。
	go func(count int) {
		for j := 1; j <= count; j++ {
			recvCond.L.Lock()
			for mailbox == "" {
				recvCond.Wait()
			}
			log.Printf("%s", mailbox)
			mailbox = ""
			recvCond.L.Unlock()
			sendCond.Broadcast() // 此处用 Broadcast，所以 [1] 处，必须用 for
			// sendCond.Signal()    // 此处用 Signal， [1] 处， 可以用 if
		}
		wg.Done()
	}(max * workers)

	wg.Wait()
}
