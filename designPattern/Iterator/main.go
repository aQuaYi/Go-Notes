// package main 是用于说明迭代器模式
package main

import "ShowYouTheGolangCode/designPattern/Iterator/bookShelf"

func main() {
	bs := bookShelf.New()
	bs.Append(bookShelf.NewBook("Around the World in 80 Days"))
	bs.Append(bookShelf.NewBook("Bible"))
	bs.Append(bookShelf.NewBook("Cinderella"))
	bs.Append(bookShelf.NewBook("Daddy Long Legs"))
	it := bs.Iterator()
	iterateDo(it)

}

func iterateDo(it bookShelf.Iterator) {
	for it.HasNext() {
		b := it.Next()
		b.Do()
	}
}
