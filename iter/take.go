package iter

// Take 返回一个迭代器，取指定数量的元素
func Take[T any](it Iterator[T], n int) Iterator[T] {
	return Generator(func() GenFunc[T] {
		limit := n
		return func() (T, bool) {
			var t T
			if limit > 0 {
				limit--
				if !it.Next() {
					return t, false
				}
				return it.Value(), true
			}

			return t, false
		}
	}())
}
