package main

import (
	"fmt"
)

//SayHier 是一个问候的接口
type SayHier interface {
	SayHi()
}

type people struct {
	SayHier              //嵌入一个接口，people可以访问接口下的所有方法
	SayBye  func(string) //嵌入一个方法，以people.SayBye(string)方式执行
}

//Chinese 是中国人structure
type Chinese struct {
	name string
}

//SayHi 是中文的问候方式
func (c Chinese) SayHi() {
	fmt.Printf("你好，我是%s\n", c.name)
}

//以中文方式告别
func cBye(name string) {
	fmt.Println("再见，", name)
}

//American 是美国人的structure
type American struct {
	name string
}

//SayHi 是美语的问候方式
func (a American) SayHi() {
	fmt.Printf("Hello, I'm %s\n", a.name)
}

//以美语方式告别
func aBye(name string) {
	fmt.Println("Bye Bye,", name)
}

func main() {
	c := Chinese{name: "张三"}
	cp := people{SayHier: c, SayBye: cBye}
	cp.SayHi()
	cp.SayBye("李四")

	a := American{name: "Tom"}
	ap := people{SayHier: a, SayBye: aBye}
	ap.SayHi()
	ap.SayBye("Jerry")
}
