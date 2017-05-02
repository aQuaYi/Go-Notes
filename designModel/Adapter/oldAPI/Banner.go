package oldAPI

import "fmt"

type Banner struct {
	str string
}

func NewBanner(str string) *Banner {
	return &Banner{str: str}
}

func (b *Banner) ShowWithParen() {
	fmt.Println("(" + b.str + ")")
}

func (b *Banner) ShowWithAster() {
	fmt.Println("*" + b.str + "*")
}
