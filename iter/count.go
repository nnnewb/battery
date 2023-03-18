package iter

// Count 统计迭代器元素总数
func Count[T any](iterator Iterator[T]) int {
	var i int
	for !iterator.Exhausted() {
		iterator = iterator.Next()
		i++
	}
	return i
}
