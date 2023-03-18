package iter

import "github.com/nnnewb/battery/optional"

type chanIter[T any] struct {
	c   <-chan optional.Optional[T]
	cur optional.Optional[T]
}

func (i chanIter[T]) Exhausted() bool {
	return !i.cur.HasValue()
}

func (i chanIter[T]) Next() Iterator[T] {
	i.cur = <-i.c
	return i
}

func (i chanIter[T]) Value() T {
	return i.cur.Value()
}

// Chan 从迭代器创建一个 go 内置 chan 对象
func Chan[T any](iterator Iterator[T]) <-chan T {
	c := make(chan T)
	go func() {
		for !iterator.Exhausted() {
			c <- iterator.Value()
			iterator = iterator.Next()
		}
	}()
	return c
}

var _ Iterator[struct{}] = chanIter[struct{}]{}
