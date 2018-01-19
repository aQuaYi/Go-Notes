package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS = ", n)
	loop(n)
}

func count() {
	x := 0
	for i := 0; i < 1<<34-1; i++ {
		x += i
	}
	fmt.Println(x)
}

func loop(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}
