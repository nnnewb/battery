package iter

// Skip 返回一个跳过n个元素的迭代器.
func Skip[T any](it Iterator[T], n int) Iterator[T] {
	return Generator(func() GenFunc[T] {
		return func() (T, bool) {
			var t T
			for it.Next() {
				n--
				if n <= 0 {
					break
				}
			}
			if !it.Next() {
				return t, false
			}
			return it.Value(), true
		}
	}())
}
