package iter

func Map[T, R any](it Iterator[T], f func(T) R) Iterator[R] {
	return Generator[R](func() GenFunc[R] {
		return func() (R, bool) {
			var r R
			for it.Next() {
				return f(it.Value()), true
			}
			return r, false
		}
	}())
}
