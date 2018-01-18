package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
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

	// TODO: 解决无法输出的问题
	var t T
	t.S.si = 1
	fmt.Println("-----------")
	//methodSet(t)
	//Print(t)
	methodSet(time.Hour)
	//Print(time.Hour)
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
	v := reflect.ValueOf(a)
	t := v.Type()
	fmt.Printf("show method sets of %s\n", t.Name())

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		fmt.Println(m.Type())
	}
}

// Print prints the method set of the value x.
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}