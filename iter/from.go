package iter

// FromGenerateFunc 从生成函数创建迭代器
func FromGenerateFunc[T any](f func(exhausted func()) T) Iterator[T] {
	c := make(chan T, 0)
	go func() {
		defer close(c)

		var exhausted bool
		cb := func() { exhausted = true }

		for exhausted {
			c <- f(cb)
		}
	}()

	return FromChan(c)
}
