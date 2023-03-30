package iter

// ToChannel 从迭代器创建一个 go 内置 chan 对象
func ToChannel[T any](iterator Iterator[T]) <-chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for iterator.Next() {
			c <- iterator.Value()
		}
	}()
	return c
}
