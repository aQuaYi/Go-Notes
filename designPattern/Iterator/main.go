// package main 是用于说明迭代器模式
package main

import (
	"github.com/aQuaYi/show-You-the-Go-Code/designPattern/Iterator/bookShelf"
)

func main() {
	bs := bookShelf.New()
	bs.Append("Around the World in 80 Days")
	bs.Append("Bible")
	bs.Append("Cinderella")
	bs.Append("Daddy Long Legs")
	bookCh := bs.Iterator()
	for b := range bookCh {
		b.Do()
	}
}
