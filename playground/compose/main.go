package main

import "fmt"

type aOne struct {
}

func (ao aOne) a() {
	fmt.Println("is aOneS")
}

type aTwo struct {
}

func (at aTwo) a() {
	fmt.Println("is aTwoS")
}

type aa struct {
	aOne
	aTwo
}

func main() {
	aaa := aa{aOne: aOne{}, aTwo: aTwo{}}
	aaa.a()
}
