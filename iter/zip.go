package iter

type zipIter[T1, T2, R any] struct {
	it1 Iterator[T1]
	it2 Iterator[T2]
	cur R
	f   func(T1, T2) R
}

func (i zipIter[T1, T2, R]) Next() Iterator[R] {
	i.it1 = i.it1.Next()
	i.it2 = i.it2.Next()
	if !i.it1.Exhausted() && i.it2.Exhausted() {
		i.cur = i.f(i.it1.Value(), i.it2.Value())
	}
	return i
}

func (i zipIter[T1, T2, R]) Value() R {
	return i.cur
}

func (i zipIter[T1, T2, R]) Exhausted() bool {
	return i.it1.Exhausted() || i.it2.Exhausted()
}

// Zip 从两个迭代器创建新的迭代器，新迭代器对 iterator 和 iterator2 的元素调函数 f 。
// 若 iterator 可迭代次数不等，则 zip 迭代器迭代次数不超过 iterator 和 iterator2 中
// 可迭代次数较短者。
func Zip[T1, T2, R any](iterator Iterator[T1], iterator2 Iterator[T2], f func(T1, T2) R) Iterator[R] {
	it := zipIter[T1, T2, R]{
		it1: iterator,
		it2: iterator2,
		f:   f,
	}
	if !it.Exhausted() {
		it.cur = it.f(it.it1.Value(), it.it2.Value())
	}
	return it
}
