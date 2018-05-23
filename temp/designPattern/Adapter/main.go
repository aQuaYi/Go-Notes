package main

import (
	"fmt"
)

func main() {
	pb := newBanner("Hello")
	pb.printWeak()
}

/*
以下是新的API
banner.PrintWeak()与oldBanner.showWithBrackets()实现的一模一样的功能，但是，必须使用PrintWeak()的名字。
*/

type banner struct {
	*oldBanner
}

func newBanner(str string) *banner {
	return &banner{&oldBanner{str: str}}
}

func (b *banner) printWeak() {
	//关键结构
	//调用printWeak的时候，转向了showWithBrackets，实现了适配
	b.showWithBrackets()
}

/*
以下是旧的API
oldBanner.ShowWithBrackets()输出banner的内容，并在前后加上括号。
*/

type oldBanner struct {
	str string
}

func (ob *oldBanner) showWithBrackets() {
	fmt.Println("(" + ob.str + ")")
}
