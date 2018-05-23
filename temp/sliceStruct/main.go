package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	const cap = 5

	stack := make([]int, 0, cap)

	push := func(x int) error {
		n := len(stack)
		if n == cap {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	pop := func() (int, error) {
		n := len(stack)
		if n == 0 {
			return 0, errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}

	// 入栈
	for i := 0; i < 7; i++ {
		fmt.Printf("push: %d, %v, %v\n", i, push(i), stack)
	}

	// 出栈
	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop : %d, %v, %v\n", x, err, stack)
	}
}

func createArrayAndAssign() [1024]int {
	var res [1024]int
	for i := 0; i < len(res); i++ {
		res[i] = i
	}
	return res
}

func createArray() [1024]int {
	var res [1024]int
	return res
}

func makeSliceAndAssign() []int {
	res := make([]int, 1024)
	for i := 0; i < len(res); i++ {
		res[i] = i
	}
	return res
}

func makeSlice() []int {
	res := make([]int, 1024)
	return res
}
