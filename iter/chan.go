package iter

type chanIter[T any] struct {
	c         <-chan T
	cur       T
	exhausted bool
}

func (i chanIter[T]) Exhausted() bool {
	return i.exhausted
}

func (i chanIter[T]) Next() Iterator[T] {
	var ok bool
	i.cur, ok = <-i.c
	i.exhausted = !ok
	return i
}

func (i chanIter[T]) Value() T {
	return i.cur
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
