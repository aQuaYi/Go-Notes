package main

import (
	"fmt"
)

type people struct {
	age int
}

func (p *people) BeBorn() {
	if p == nil {
		//下方的赋值语句是无效的。
		//当p!=nil时，可以修改p.age
		//但是，不能把p自身，即不能改变p==nil的事实。
		p = &people{
			age: 0,
		}
		return
	}
}

func main() {
	var nilPeople *people
	nilPeople.BeBorn()
	fmt.Println("在BeBorn后，nilPeople =", nilPeople)
}
