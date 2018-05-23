package main

import (
	"fmt"
	"reflect"
)

type X int

func main() {
	var a X = 100
	t := reflect.TypeOf(a)

	fmt.Println(t.Name(), t.Kind())
}
