package iter

type GenFunc[T any] func() (T, bool)

type genIter[T any] struct {
	f   func() (T, bool)
	cur T
}

func (i *genIter[T]) Next() bool {
	var ok bool
	i.cur, ok = i.f()
	return ok
}

func (i *genIter[T]) Value() T {
	return i.cur
}

// Generator 从生成函数创建迭代器
func Generator[T any](f GenFunc[T]) Iterator[T] {
	return &genIter[T]{
		f: f,
	}
}
