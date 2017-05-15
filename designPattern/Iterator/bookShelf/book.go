package bookShelf

import "fmt"

//Booker
type Booker interface {
	Do()
}

// Book 定义了书籍类
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
