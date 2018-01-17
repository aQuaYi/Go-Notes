package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4}
	p := &a

	// 数组指针可以直接用来操作元素
	p[1] += 10
	fmt.Println(p)

	s := a[:]
	sp := &s
	// sp[1]+=10
	// 切片的指针无法直接操作元素
	(*sp)[1] += 10
	fmt.Println(sp)
}
