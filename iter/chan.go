package iter

// FromChan 从 go 内置的 chan 创建 Iterator.
func FromChan[T any](c <-chan T) Iterator[T] {
	return Generator(func() GenFunc[T] {
		return func() (T, bool) {
			v, ok := <-c
			return v, ok
		}
	}())
}
