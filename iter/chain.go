package iter

// Chain 链接多个迭代器，返回新的迭代器。新的迭代器按参数传入顺序迭代返回。
// 当所有迭代器都穷尽，则本迭代器也视为穷尽。
func Chain[T any](iterators ...Iterator[T]) Iterator[T] {
	return Generator(func() GenFunc[T] {
		var idx int
		return func() (T, bool) {
			var t T
			if len(iterators) == 0 {
				return t, false
			}

			if idx >= len(iterators) {
				return t, false
			}

			for {
				ok := iterators[idx].Next()
				if !ok {
					idx++
					if idx >= len(iterators) {
						return t, false
					}
					continue
				}
				return iterators[idx].Value(), ok
			}
		}
	}())
}
