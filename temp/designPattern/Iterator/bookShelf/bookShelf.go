package bookShelf

//BookShelfer 汇集了书架所需的方法
type BookShelfer interface {
	Append(string)
	Get(int) Booker
	Len() int
	Iterator() <-chan Booker
}

//bookShelf 定义了书架类
type bookShelf struct {
	books  []Booker
	number int
}

//New 返回了一个BookShelfer接口
func New() BookShelfer {
	return &bookShelf{}
}

func (bs *bookShelf) Append(name string) {
	bs.books = append(bs.books, newBook(name))
	bs.number++
}

func (bs *bookShelf) Get(i int) Booker {
	return bs.books[i]
}

func (bs *bookShelf) Len() int {
	return bs.number
}

func (bs *bookShelf) Iterator() <-chan Booker {
	bookCh := make(chan Booker)

	go func() {
		defer close(bookCh)
		for _, v := range bs.books {
			bookCh <- v
		}
	}()

	return bookCh
}
