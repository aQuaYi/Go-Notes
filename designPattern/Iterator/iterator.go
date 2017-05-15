// package main 是用于说明迭代器模式
package main

import (
	it "ShowYouTheGolangCode/designPattern/Iterator/Iterator"
)

func main() {
	bs := &it.BookShelf{}
	bs.AppendBook(&it.Book{Name: "Around the World in 80 Days"})
	bs.AppendBook(&it.Book{Name: "Bible"})
	bs.AppendBook(&it.Book{Name: "Cinderella"})
	bs.AppendBook(&it.Book{Name: "Daddy Long Legs"})
	it := bs.Iterator()
	doSomething(it)
}

func doSomething(it it.Iterator) {
	for {
		if !it.HasNext() {
			return
		}
		b := it.Next()
		b.Do()
	}
}
