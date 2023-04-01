package iter

// First 返回首个符合条件的元素的迭代器.
func First[T any](it Iterator[T], predicate func(T) bool) (T, bool) {
	var zero T
	i := Filter(it, predicate)
	for i.Next() {
		return i.Value(), true
	}
	return zero, false
}
