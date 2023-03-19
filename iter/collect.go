package iter

// Collect 收集迭代器元素，返回一个切片
func Collect[T any](iterator Iterator[T]) []T {
	ret := make([]T, 0)
	for iterator = iterator.Next(); !iterator.Exhausted(); iterator = iterator.Next() {
		ret = append(ret, iterator.Value())
	}
	return ret
}
