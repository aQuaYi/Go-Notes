package templateMethod

import "fmt"

// Display 提供了Display()方法。
type Display interface {
	Display()
}

type displayAble interface {
	open()
	print()
	close()
}

type abstractDisplay struct {
	inst displayAble
}

func (ad *abstractDisplay) Display() {
	if ad.inst == nil {
		return
	}

	ad.inst.open()
	for i := 0; i < 5; i++ {
		ad.inst.print()
	}
	ad.inst.close()
}

type charDisplay struct {
	ch byte
}

// NewCharDisplay 返回charDisplay的Display接口的形式
func NewCharDisplay(b byte) Display {
	result := &abstractDisplay{}
	result.inst = &charDisplay{ch: b}
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

// NewStringDisplay 返回stringDisplay的Display接口形式
func NewStringDisplay(s string) Display {
	result := &abstractDisplay{}
	result.inst = &stringDisplay{str: s}
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
