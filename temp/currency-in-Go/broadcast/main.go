package main

import (
	"fmt"
	"sync"
)

type button struct {
	// call
	clicked *sync.Cond
	// back
	mu          *sync.Mutex
	actionCount int
	actions     *sync.WaitGroup
}

func (b *button) click() {
	b.actions.Add(b.actionCount)
	b.clicked.Broadcast()
	b.actions.Wait()
}

func (b *button) subscribe(fn func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		for {
			b.clicked.L.Lock()
			b.clicked.Wait()
			fn()
			b.clicked.L.Unlock()
			b.actions.Done()
		}
	}()
	wg.Wait()
	b.mu.Lock()
	b.actionCount++
	b.mu.Unlock()
}

func main() {
	btn := &button{
		mu:      &sync.Mutex{},
		clicked: sync.NewCond(&sync.Mutex{}),
		actions: &sync.WaitGroup{},
	}

	btn.subscribe(func() {
		fmt.Println("Maximizing window.")
	})
	btn.subscribe(func() {
		fmt.Println("Displaying annoying dialog box!")
	})
	btn.subscribe(func() {
		fmt.Println("Mouse clicked.")
	})

	s := say{word: "hello"}
	btn.subscribe(s.hello)

	clickTimes := 5
	for i := 0; i < clickTimes; i++ {
		fmt.Println("---new---click---")
		btn.click()
	}

}

type say struct {
	word string
}

func (s say) hello() {
	fmt.Printf("say %s\n", s.word)
}
