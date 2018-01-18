package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n num = 25

	// 方法是特殊的函数，
	// 方法有一个隐藏的参数，就是 receiver
	fmt.Printf("n: %p, %v\n", &n, n)
	n.value()
	n.pointer()
	fmt.Printf("n: %p, %v\n", &n, n)

	// 基础类型的实例方法和指针方法可以相互调用
	// 调用时，按照方法的参数类型进行转换
	fmt.Println("-----------")
	p := &n

	n.value()
	n.pointer()

	p.value()
	p.pointer()

	var t T
	fmt.Println("--查看 T 的方法集------------------------")
	methodSet(t)
	fmt.Println("--查看 *T 的方法集-----------------------")
	methodSet(&t)
	fmt.Println("--SPtr 不在 T 的方法集中, 依然可以被调用---")
	fmt.Println("--call t.SPtr()---")
	t.SPtr()
}

type num int

func (n num) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *num) pointer() {
	*n++
	fmt.Printf("p: %p, %v\n", n, *n)
}

type S struct {
}

type T struct {
	S
}

func (S) SVal() {}

func (*S) SPtr() {
	fmt.Println("I'm sPtr")
}

func (T) TVal() {}

func (*T) TPtr() {}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		println(m.Name)
	}
}

