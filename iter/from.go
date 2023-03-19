package iter

// FromSlice 创建一个切片迭代器
func FromSlice[T any](s []T) Iterator[T] {
	return sliceIter[T]{
		underlying: s,
		idx:        -1,
	}
}

// FromChan 从 go 内置的 chan 创建 Iterator
func FromChan[T any](c <-chan T) Iterator[T] {
	return chanIter[T]{c: c}
}

// FromGenerateFunc 从生成函数创建迭代器
func FromGenerateFunc[T any](f func(exhausted func()) T) Iterator[T] {
	return FromChan(gen(f))
}
