package bookShelf

import "fmt"

//Booker 包含了book所有的暴露方法
type Booker interface {
	Do()
}

type book struct {
	name string
}

//newBook 创建一本新书
func newBook(name string) Booker {
	return &book{name: name}
}

// Do 让书本读出自己的名字。
func (b *book) Do() {
	fmt.Printf("My name is %s\n", b.name)
}
