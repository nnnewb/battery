package iter

// Zip 从两个迭代器创建新的迭代器，新迭代器对 iterator 和 iterator2 的元素调函数 f 。
// 若 iterator 可迭代次数不等，则 zip 迭代器迭代次数不超过 iterator 和 iterator2 中
// 可迭代次数较短者。
func Zip[T1, T2, R any](iterator Iterator[T1], iterator2 Iterator[T2], f func(T1, T2) R) Iterator[R] {
	return Generator(func() GenFunc[R] {
		return func() (R, bool) {
			var t R
			if !iterator.Next() {
				return t, false
			}

			if !iterator2.Next() {
				return t, false
			}

			return f(iterator.Value(), iterator2.Value()), true
		}
	}())
}
