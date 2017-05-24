package main

import (
	"ShowYouTheGolangCode/Golang/unexporting/unexportedType/counters"
	"fmt"
)

func main() {

	//通过返回值，变量counter被赋值为counters.alertCounter类型的值。
	counter := counters.New(10)
	fmt.Printf("counter's type is %T", counter)

	//undefined: counerts in counerts.alertCounter
	counter2 := counerts.alertCounter(10)
}
