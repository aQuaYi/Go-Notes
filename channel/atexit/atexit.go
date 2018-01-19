package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// 能够在程序退出时，完成指定的工作
// 无论是正常退出，还是 panic 的
func main() {
	atExit(func() { fmt.Println("exit func 1 ....") })
	atExit(func() { fmt.Println("exit func 2 ....") })

	waitExit()
}

var exits = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atExit(f func()) {
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs, f)
}

func waitExit() {
	fmt.Println("正在等待程序退出，请按 Ctrl+C")

	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}

	exits.RLock()
	for _, f := range exits.funcs {
		//noinspection GoDeferInLoop
		defer f()
		// defer 确保 panic 了以后，也能够执行 f()
		// funcs 按照 FILO 顺序执行
	}
	exits.RUnlock()

	<-exits.signals
}
