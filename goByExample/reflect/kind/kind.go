package main

import (
	"fmt"
	"reflect"
)

// X int的X变种
type X int

// Y X的Y变种
type Y X

func main() {
	var x1, x2 X = 1, 2
	var y3 Y = 3

	// reflect.TypeOf 返回的是变量的静态类型，就是定义变量时，选择的类型。
	tx1, tx2, ty3 := reflect.TypeOf(x1), reflect.TypeOf(x2), reflect.TypeOf(y3)
	fmt.Println(tx1 == tx2, tx1 == ty3)

	// reflect.Type.Kind() 返回的是，变量的底层类型。
	fmt.Println(tx1.Kind() == ty3.Kind())
}
