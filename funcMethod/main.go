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

	//
	var t T
	t.S.si = 1
	fmt.Println("-----------")
	methodSet(t)
	fmt.Println("-----------")
	methodSet(&t)
	fmt.Println("-----------")
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
	si int
}

type T struct {
	S
}

func (s S) sVal() {}

func (s *S) sPtr() {}

func (t T) tVal() {}

func (t *T) tPtr() {}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("show method sets of %s\n", t.Name())

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
