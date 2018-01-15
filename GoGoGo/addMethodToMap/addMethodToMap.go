// myMap.go
// this is my map
package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	mm := make(myMap)
	mm.Set("a", "1")
	mm.Set("b", "2")
	mm.Set("c", "3")
	fmt.Println(mm)
	mm.Del("b")
	fmt.Println(mm)
}

type myMap map[string]string

func (m myMap) Set(key, value string) {
	mu.Lock()
	m[key] = value
	mu.Unlock()
}
func (m myMap) Get(key string) string {
	return m[key]
}

func (m myMap) Del(key string) {
	mu.Lock()
	delete(m, key)
	mu.Unlock()
}

type Storage interface {
	Del(key string)
	Get(key string) string
	Set(key, value string)
}
