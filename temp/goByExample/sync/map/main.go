package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	number := 5
	list := make(map[int]int)

	for i := 0; i < number; i++ {
		list[i] = i
	}

	var m sync.Map
	// 把 list 的键值对，存入m
	for k, v := range list {
		m.Store(k, v)
	}

	zero, ok := m.Load(0)
	fmt.Printf("zero, ok := m.Load(0), zero = %d, ok = %t\n", zero, ok)
	noKey, ok := m.Load(number)
	fmt.Printf("noKey, ok := m.Load(%d), noKey = %v, ok = %t\n", number, noKey, ok)

	zero, loaded := m.LoadOrStore(0, "zero")
	fmt.Printf(`zero, loaded := m.LoadOrStore(0, "zero") -->> zero = %v, loaded = %t`, zero, loaded)
	fmt.Println()

	noKey, loaded = m.LoadOrStore(number, "number")
	fmt.Printf(`noKey, loaded = m.LoadOrStore(number, "number") -->> noKey = %v, loaded = %t`, noKey, loaded)
	fmt.Println()

	m.Delete(number)
	fmt.Printf("已经删除了 key = %d\n", number)

	changeDuringRange(m, number)

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		time.Sleep(time.Second)
		return true
	})
	fmt.Println("可以看到，Range 输出的value，总是最新的值。")
}

func changeDuringRange(m sync.Map, number int) {
	go func() {
		time.Sleep(500 * time.Millisecond)
		base := 1
		for i := 0; i < number-1; i++ {
			base = base * 10
			for j := 0; j < number; j++ {
				m.Store(j, j+base)
			}
			time.Sleep(time.Second)
		}
	}()
}
