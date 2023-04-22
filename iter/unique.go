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
				// If the key is already in the map, skip it
			}
			return zero, false
		}
	}())
}

// UniqueBy 对迭代器元素进行去重，使用 key 参数指定去重依据。
func UniqueBy[T, K comparable](it Iterator[T], key func(T) K) Iterator[T] {
	return Generator(func() GenFunc[T] {
		s := make(map[K]bool)
		return func() (T, bool) {
			var zero T
			for it.Next() {
				k := key(it.Value())
				if !s[k] {
					s[k] = true
					return it.Value(), true
				}
				// If the key is already in the map, skip it
			}
			return zero, false
		}
	}())
}
