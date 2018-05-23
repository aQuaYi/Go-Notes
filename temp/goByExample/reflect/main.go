package main

import (
	"fmt"
	"reflect"
)

type myInt int

func main() {
	var mi myInt = 1
	t := reflect.TypeOf(mi)

	fmt.Println(t.Name(), t.Kind())
}
