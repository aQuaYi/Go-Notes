package main

var (
	maxWorker = 4
	maxQueue  = 4
)

func main() {
	d := newDispatcher(maxWorker)
	d.run()
	loadMaker()
}
