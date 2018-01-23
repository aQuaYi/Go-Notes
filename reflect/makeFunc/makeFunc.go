package main

import (
	"fmt"
	"reflect"
	"strings"
)

func add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value

	switch args[0].Kind() {
	case reflect.Int:
		n := 0

		for _, a := range args {
			n += int(a.Int())
		}

		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, 0, len(args))

		for _, s := range args {
			ss = append(ss, s.String())
		}

		ret = reflect.ValueOf(strings.Join(ss, ""))
	}
	results = append(results, ret)
	return
}

func makeAdd(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), add)
	fn.Set(v)
}

func main() {
	var intAdd func(x, y int) int
	var stringAdd func(a, b string) string

	makeAdd(&intAdd)
	makeAdd(&stringAdd)

	fmt.Println(intAdd(1, 2))
	fmt.Println(stringAdd("hello, ", "world!"))
}
