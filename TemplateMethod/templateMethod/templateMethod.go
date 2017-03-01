package templateMethod

import (
	"fmt"
)

type IDisplay interface {
	Open()
	Print()
	Close()
}

type AbstractDisplay struct {
	Inst IDisplay
}

func (ad *AbstractDisplay) Display() {
	if ad.Inst == nil {
		return
	}

	ad.Inst.Open()
	for i := 0; i < 5; i++ {
		ad.Inst.Print()
	}
	ad.Inst.Close()
}

type CharDisplay struct {
	Ch byte
}

func (cd *CharDisplay) Open() {
	fmt.Print("<<")
}

func (cd *CharDisplay) Print() {
	fmt.Print(string(cd.Ch))
}

func (cd *CharDisplay) Close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	Str string
}

func (sd *StringDisplay) printLine() {
	length := len(sd.Str)
	fmt.Print("+")
	for i := 0; i < length; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (sd *StringDisplay) Open() {
	sd.printLine()
}

func (sd *StringDisplay) Print() {
	fmt.Println("|" + sd.Str + "|")
}

func (sd *StringDisplay) Close() {
	sd.printLine()
}
