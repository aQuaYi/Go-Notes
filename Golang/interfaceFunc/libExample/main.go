package main

import (
	"ShowYouTheGolangCode/Golang/interfaceFunc/libExample/lib"
	"fmt"
)

func selfAge(k, v interface{}) {
	fmt.Printf("大家好，我叫%s, 今年%d岁\n", k, v)
}

func selfNumber(k, v interface{}) {
	fmt.Printf("大家好，我叫%s, 学号是%d\n", k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	lib.Each(persons, selfAge)    //selfAge满足lib.HandleFunc类型，不必转换
	lib.Each(persons, selfNumber) //selfNumber满足lib.HandleFunc类型，不必转换

	i := 1
	lib.TestMyInt(lib.MyInt(i)) //int不能满足lib.MyInt类型，需要转换
	//lib.TestMyInt(i)
}
