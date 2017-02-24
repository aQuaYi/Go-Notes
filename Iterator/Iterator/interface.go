package iterator

type Worker interface {
	Do()
}
type Iterator interface {
	HasNext() bool
	Next() Worker
}

type Aggregate interface {
	Iterator() Iterator
}
