package main

import (
	"runtime"
	"time"
)

func main() {
	workers := runtime.NumCPU() / 2

	for i := 0; i < workers; i++ {
		go func() {
			for {
			}
		}()
	}

	time.Sleep(60 * time.Second)
}
