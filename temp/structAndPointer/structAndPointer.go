package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	d := newDog("no name")
	dp := &d

	// 对象可以调用其指针的方法
	d.SetName("doggy") // 相当于 (&d).SetName("monster")
	fmt.Println(d.Name)

	// 指针也可以调用其对象的方法
	dp.Barking()

	// 指针的方法集，包含了其对象的方法，可以满足指针的要求
	dpi := doger(dp)
	dpi.Barking()

	// 对象的方法集，不包含其指针的方法，不能满足指针的要求
	// di := Doger(d)
	// di.Barking()

}

type dog struct {
	Name string
}

func (d *dog) SetName(name string) {
	d.Name = name
}

func (d dog) Barking() {
	fmt.Println("wow wow")
}

func newDog(name string) dog {
	return dog{
		Name: name,
	}
}

type doger interface {
	SetName(string)
	Barking()
}
