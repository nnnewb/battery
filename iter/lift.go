package iter

// FromSlice 创建一个切片迭代器
func FromSlice[T any](s []T) Iterator[T] {
	return Generator[T](func() GenFunc[T] {
		var i int
		return func() (T, bool) {
			var t T
			if i >= len(s) {
				return t, false
			}
			r := s[i]
			i++
			return r, true
		}
	}())
}

// Lift 兼容 go-functional 的写法，等同于 FromSlice
func Lift[T any](s []T) Iterator[T] {
	return FromSlice(s)
}
