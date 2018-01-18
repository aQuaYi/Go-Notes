package main

import "fmt"

func main() {
	var n N = 25

	fmt.Printf("main.n: %p, %d\n", &n, n)

	// 通过被类型引用的方法，被还原成了普通函数
	// receiver 要被当做第一个参数传入
	f1 := N.test
	f1(n)

	f2 := (*N).test
	f2(&n)
}

type N int

func (n N) test() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}
