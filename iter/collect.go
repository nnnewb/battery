package iter

// Collect 收集迭代器元素，返回一个切片
func Collect[T any](iterator Iterator[T]) []T {
	ret := make([]T, 0)
	for !iterator.Exhausted() {
		ret = append(ret, iterator.Value())
		iterator = iterator.Next()
	}
	return ret
}
