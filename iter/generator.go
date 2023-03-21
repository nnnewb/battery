package iter

type GenFunc[T any] func() (T, bool)

type genIter[T any] struct {
	f         func() (T, bool)
	cur       T
	exhausted bool
}

func (i genIter[T]) Next() Iterator[T] {
	var ok bool
	i.cur, ok = i.f()
	i.exhausted = !ok
	return i
}

func (i genIter[T]) Value() T {
	return i.cur
}

func (i genIter[T]) Exhausted() bool {
	return i.exhausted
}

// FromGenerateFunc 是 Generator 的别名，从生成函数创建迭代器
func FromGenerateFunc[T any](f GenFunc[T]) Iterator[T] {
	return Generator(f)
}

// Generator 从生成函数创建迭代器
func Generator[T any](f GenFunc[T]) Iterator[T] {
	return genIter[T]{
		f:         f,
		exhausted: false,
	}
}
