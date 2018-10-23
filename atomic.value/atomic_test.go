package noway_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

var mu sync.RWMutex

type manager struct {
	agents int
}

func BenchmarkManagerLock(b *testing.B) {
	i := 0
	m := new(manager)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			m.agents = i
			mu.Unlock()
		}
	})
}

func BenchmarkManagerRLock(b *testing.B) {
	m := manager{agents: 100}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.RLock()
			_ = m.agents
			mu.RUnlock()
		}
	})
}

func BenchmarkManagerAtomicValueStore(b *testing.B) {
	var managerVal atomic.Value
	m := manager{agents: 0}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			managerVal.Store(m)
		}
	})
}

func BenchmarkManagerAtomicValueLoad(b *testing.B) {
	var managerVal atomic.Value
	managerVal.Store(&manager{agents: 100})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = managerVal.Load().(*manager)
		}
	})
}
