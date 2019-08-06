package orderprint

import (
	"runtime"
	"sync"
)

// Worker  工作队列里的工作者
type Worker interface {
	Work()
}

func newRing(workers ...Worker) *Ring {
	return &Ring{
		workers: workers,
	}
}

// Ring 环状队列
type Ring struct {
	workers []Worker
	pos     int
}

type holdOnWorker struct {
	count    int
	trigger  sync.Mutex
	finished sync.Mutex
}

func newHoldOnWorker() *holdOnWorker {
	w := &holdOnWorker{}
	w.trigger.Lock()
	w.finished.Lock()
	return w
}

// Work 触发事件
func (w *holdOnWorker) Work() {
	w.trigger.Unlock()
	w.finished.Lock()
}

// HoldOn 接受事件
func (w *holdOnWorker) HoldOn() {
	w.trigger.Lock()
	w.count++
	w.finished.Unlock()
}

func mutex() {
	w1 := newHoldOnWorker()
	go fakeThread("1", w1.HoldOn)

	w2 := newHoldOnWorker()
	go fakeThread("2", w2.HoldOn)

	w3 := newHoldOnWorker()
	go fakeThread("3", w3.HoldOn)

	r := newRing(w1, w2, w3)
	for i := 0; i < 300; i++ {
		if r.pos == len(r.workers) {
			r.pos = 0
		}
		r.workers[r.pos].Work()
		r.pos++
	}
}

// fakeThread 假的工作线程
func fakeThread(id string, holdOn func()) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for {
		holdOn()
	}
}
