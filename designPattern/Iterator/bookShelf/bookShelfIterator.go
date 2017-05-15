package bookShelf

//Iterator 具有迭代功能
type Iterator interface {
	HasNext() bool
	Next() Booker
}

type bookShelfIterator struct {
	bookShelf *bookShelf
	index     int
}

func (bsi *bookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.Len()
}

func (bsi *bookShelfIterator) Next() Booker {
	b := bsi.bookShelf.Get(bsi.index)
	bsi.index++
	return b
}
