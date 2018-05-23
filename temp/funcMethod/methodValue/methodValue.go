package main

import "fmt"

func main() {
	var n N = 100
	p := &n

	// 被实例或指针引用的方法，被称为 method value
	// method value 在被复制时，会复制 receiver 的值，连同 method value 一起保存
	// 以便在稍后的调用中，把 receiver 的值传入 method

	fmt.Println("-----------------")
	fmt.Printf("main.n     : %d, %p\n", n, p)

	n++
	fv1 := n.vTest
	// vTest 的 receiver 的类型是 N
	// fv1 中保存了从 n 复制的 101

	n++
	fv2 := p.vTest
	// 虽然此时是通过指针调用的 vTest
	// 但是，vTest 的 receiver 的类型是 N
	// fv2 中保存了从 *p 复制的 102

	n++
	fmt.Printf("main.n     : %d, %p\n", n, p)

	fv1("fv1")
	fv2("fv2")

	fmt.Println("--以上，被复制的是变量--")
	fmt.Println("-----------------")
	fmt.Printf("main.n     : %d, %p\n", n, p)

	n++
	fp1 := n.pTest
	// 虽然此时是通过实例，调用的 pTest
	// 但是，pTest 的 receiver 的类型是 *N
	// fp1 中保存了从 &n 复制了 n 的指针

	n++
	fp2 := p.pTest
	// pTest 的 receiver 的类型是 *N
	// fp2 中保存了从 p 复制了 n 的指针

	n++
	fmt.Printf("main.n     : %d, %p\n", n, p)

	fp1("fp1")
	fp2("fp2")
	fmt.Println("--以上，被复制的是指针--")
}

type N int

func (n N) vTest(name string) {
	fmt.Printf("%s.vTest.n: %d, %p\n", name, n, &n)
}

func (n *N) pTest(name string) {
	fmt.Printf("%s.pTest.n: %d, %p\n", name, *n, n)
}
