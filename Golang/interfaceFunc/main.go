package main

import (
	"fmt"
)

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func each(m map[interface{}]interface{}, h Handler) {
	for k, v := range m {
		h.Do(k, v)
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	each(m, HandlerFunc(f))
}

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

	EachFunc(persons, selfAge)
	EachFunc(persons, selfNumber)
}
