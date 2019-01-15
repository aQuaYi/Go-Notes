package main

import (
	"fmt"
	"sync"
)

// 一个 goroutine 只能输出 "ABCD..."
// 一个 goroutine 只能输出 "1234..."
// 最后输出的结果是 "A1B2C3D4E5F6G7H8I9"

func main() {
	var number sync.WaitGroup
	number.Add(1)
	var letter sync.WaitGroup
	letter.Add(1)

	var all sync.WaitGroup
	all.Add(2)

	// print numbers
	go func() {
		for i := 1; i < 10; i++ {
			letter.Wait()
			fmt.Print(i)
			number.Done()
			letter.Add(1)
		}
		all.Done()
	}()

	// print letters
	go func() {
		for i := 'A'; i < 'A'+9; i++ {
			number.Wait()
			fmt.Print(string(byte(i)))
			letter.Done()
			number.Add(1)
		}
		all.Done()
	}()

	number.Done()

	all.Wait()
}
