package iter

// Last 返回最后一个符合条件的元素的迭代器
func Last[T any](it Iterator[T], predicate func(T) bool) (T, bool) {
	var (
		last  T
		found bool
	)
	for it.Next() {
		if predicate(it.Value()) {
			last = it.Value()
			found = true
		}
	}
	return last, found
}
