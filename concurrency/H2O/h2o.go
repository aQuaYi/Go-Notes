package water

import (
	"context"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

// H2O is water
type H2O struct {
	semaH   *semaphore.Weighted
	semaO   *semaphore.Weighted
	barrier cyclicbarrier.CyclicBarrier
}

// NewH2O returns H2O pointer
func NewH2O() *H2O {
	return &H2O{
		semaH:   semaphore.NewWeighted(2),
		semaO:   semaphore.NewWeighted(1),
		barrier: cyclicbarrier.New(3),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	// Acquire 信号量空闲，以便获取。
	// 能够获取到，说明上一次的组建工作已经完成了。
	h2o.semaH.Acquire(context.Background(), 1)

	// 真正的制造工作，把 "H" 输出到需要的地方。
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()

	// barrier.Await 作为同步点
	// 上下分别属于两个不同的 h2o 组建过程
	h2o.barrier.Await(context.Background())

	// 释放信号量，为下一次组建提供机会
	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)
	// releaseOxygen() outputs "O". Do not change or remove this line.
	releaseOxygen()
	h2o.barrier.Await(context.Background())
	h2o.semaO.Release(1)
}
