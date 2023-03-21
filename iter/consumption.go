package iter

// ToChannel 从迭代器创建一个 go 内置 chan 对象
func ToChannel[T any](iterator Iterator[T]) <-chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for iterator = iterator.Next(); !iterator.Exhausted(); iterator = iterator.Next() {
			c <- iterator.Value()
		}
	}()
	return c
}

// Collect 收集迭代器元素，返回一个切片
func Collect[T any](iterator Iterator[T]) []T {
	ret := make([]T, 0)
	for iterator = iterator.Next(); !iterator.Exhausted(); iterator = iterator.Next() {
		ret = append(ret, iterator.Value())
	}
	return ret
}
