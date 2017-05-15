package main

import "fmt"

func main() {
	c := NewChar('H')
	c.Display()

	s := NewString("Hello, World!")
	s.Display()
}

// Displayer 提供了Display()方法。
type Displayer interface {
	Display()
}

type displayAbler interface {
	open()
	print()
	close()
}

type display struct {
	displayAbler
}

func (d *display) Display() {
	if d.displayAbler == nil {
		return
	}

	d.open()
	for i := 0; i < 5; i++ {
		d.print()
	}

	d.close()
}

type charDisplay struct {
	ch byte
}

// NewChar 返回charDisplay的Display接口的形式
func NewChar(b byte) Displayer {
	result := &display{}
	result.displayAbler = &charDisplay{ch: b}
	return result
}

func (cd *charDisplay) open() {
	fmt.Print("<<")
}

func (cd *charDisplay) print() {
	fmt.Print(string(cd.ch))
}

func (cd *charDisplay) close() {
	fmt.Println(">>")
}

type stringDisplay struct {
	str string
}

// NewString 返回stringDisplay的Display接口形式
func NewString(s string) Displayer {
	result := &display{}
	result.displayAbler = &stringDisplay{str: s}
	return result
}
func (sd *stringDisplay) printLine() {
	length := len(sd.str)

	fmt.Print("+")
	for i := 0; i < length; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (sd *stringDisplay) open() {
	sd.printLine()
}

func (sd *stringDisplay) print() {
	fmt.Println("|" + sd.str + "|")
}

func (sd *stringDisplay) close() {
	sd.printLine()
}
