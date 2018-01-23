package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem() // 只有基类型才能遍历字段
	}

	// 遍历 manager 的完整结构
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}

	fmt.Println("==按名称查找字段==")
	name, _ := t.FieldByName("name")
	fmt.Println(name.Name, name.Type)

	fmt.Println("==按索引查找字段==")
	age := t.FieldByIndex([]int{0, 1})
	fmt.Println(age.Name, age.Type)
}
