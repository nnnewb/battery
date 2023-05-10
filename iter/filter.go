package iter

// Filter 创建一个迭代器，过滤出符合条件的元素。
func Filter[T any](iterator Iterator[T], predicate func(T) bool) Iterator[T] {
	return Generator(func() GenFunc[T] {
		return func() (T, bool) {
			var t T
			for iterator.Next() {
				if predicate(iterator.Value()) {
					return iterator.Value(), true
				}
			}
			return t, false
		}
	}())
}
