package main

import (
	"fmt"
	"reflect"
)

type A int

type B struct {
	A
}

func (A) Av()  {}
func (*A) Ap() {}

func (B) Bv()  {}
func (*B) Bp() {}

func main() {
	var b B

	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}

	for _, t := range s {
		fmt.Println(t, ":")

		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println("  ", t.Method(i))
		}
	}
}
