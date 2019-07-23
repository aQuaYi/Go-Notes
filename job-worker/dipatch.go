package main

// dispatcher is ..
type dispatcher struct {
	size       int
	WorkerPool chan chan int
}

func newDispatcher(maxWorkers int) *dispatcher {
	pool := make(chan chan int, maxWorkers)
	return &dispatcher{
		size:       maxWorkers,
		WorkerPool: pool,
	}
}

func (d *dispatcher) run() {
	// starting n number of workers
	for i := 0; i < d.size; i++ {
		worker := newWorker(i, d.WorkerPool)
		worker.start()
	}
	go d.dispatch()
}

func (d *dispatcher) dispatch() {
	for {
		select {
		case num := <-intStream:
			// a job request has been received
			go func(num int) {
				worker := <-d.WorkerPool
				worker <- num
			}(num)
		}
	}
}
