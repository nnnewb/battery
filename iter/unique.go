package iter

// Unique 对迭代器元素进行去重。迭代器元素必须是可比较的。
func Unique[T comparable](it Iterator[T]) Iterator[T] {
	return Generator(func() GenFunc[T] {
		s := make(map[T]bool)
		return func() (T, bool) {
			var zero T
			for it.Next() {
				if !s[it.Value()] {
					s[it.Value()] = true
					return it.Value(), true
				}
			}
			return zero, false
		}
	}())
}
