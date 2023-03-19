package iter

type chainIter[T any] struct {
	iterators []Iterator[T]
	it        Iterator[T]
	idx       int
}

func (i chainIter[T]) Next() Iterator[T] {
	if len(i.iterators) == 0 {
		return i
	}

	if i.it != nil {
		i.it = i.it.Next()
	}

	// 当前迭代器已经被穷尽或未设置，更换下一个迭代器
	for i.it == nil || i.it.Exhausted() {
		if i.idx+1 >= len(i.iterators) {
			return i
		}

		i.idx++
		i.it = i.iterators[i.idx].Next()
		if !i.it.Exhausted() {
			break
		}
	}

	return i
}

func (i chainIter[T]) Value() T {
	var t T
	if i.it != nil && !i.it.Exhausted() {
		return i.it.Value()
	}
	return t
}

func (i chainIter[T]) Exhausted() bool {
	return len(i.iterators) == 0 || (i.it.Exhausted() && i.idx+1 >= len(i.iterators))
}

// Chain 链接多个迭代器，返回新的迭代器。新的迭代器按参数传入顺序迭代返回。
// 当所有迭代器都穷尽，则本迭代器也视为穷尽。
func Chain[T any](iterators ...Iterator[T]) Iterator[T] {
	return chainIter[T]{
		iterators: iterators,
		it:        nil,
		idx:       -1,
	}
}
