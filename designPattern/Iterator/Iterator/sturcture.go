package iterator

import (
	"fmt"
)

// Book 定义了书籍类
type Book struct {
	Name string
}

// Do 让书本读出自己的名字。
func (b *Book) Do() {
	fmt.Printf("My name is %s\n", b.Name)
}

// BookShelf 定义了书架类
type BookShelf struct {
	books []*Book
	last  int
}

func (bs *BookShelf) GetBookAt(i int) *Book {
	return bs.books[i]
}

func (bs *BookShelf) AppendBook(b *Book) {
	bs.books = append(bs.books, b)
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	return bs.last
}

func (bs *BookShelf) Iterator() *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bs}
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.GetLength()
}

func (bsi *BookShelfIterator) Next() Worker {
	b := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return b
}
