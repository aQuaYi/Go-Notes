package lib

import (
	"fmt"
)

type HandleFunc func(k, v interface{})

func Each(m map[interface{}]interface{}, h HandleFunc) {
	for k, v := range m {
		h(k, v)
	}
}

type MyInt int

func TestMyInt(m MyInt) {
	fmt.Println(m)
}
