package deferBenchmark

import "sync"

var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	defer m.Unlock()
}
