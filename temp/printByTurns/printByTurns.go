package main

import (
	"fmt"
	"sync"
)

// 一个 goroutine 只能输出 "ABCDEFGHI"
// 一个 goroutine 只能输出 "123456789"
// 最后输出的结果是 "A1B2C3D4E5F6G7H8I9"

func main() {
	var number sync.WaitGroup
	var letter sync.WaitGroup
	letter.Add(1)

	// print numbers
	go func() {
		for i := 1; i < 10; i++ {
			letter.Wait()
			fmt.Print(i)
			number.Done()
			letter.Add(1)
		}
	}()

	// print letters
	for i := 'A'; i < 'A'+9; i++ {
		number.Wait()
		fmt.Print(string(byte(i)))
		letter.Done()
		number.Add(1)
	}

	number.Wait()
}
