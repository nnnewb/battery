package iter

// FromChan 从 go 内置的 chan 创建 Iterator. 此函数不会关闭 chan，可以通过关闭 chan 来表示数据源已穷尽。
func FromChan[T any](c <-chan T) Iterator[T] {
	return Generator(func() GenFunc[T] {
		return func() (T, bool) {
			v, ok := <-c
			return v, ok
		}
	}())
}

// ToChan 将迭代器转换为 chan，注意此函数将取得迭代器对象所有权，在调用后请不要再使用此迭代器。
func ToChan[T any](it Iterator[T]) <-chan T {
	c := make(chan T)
	go func() {
		for it.Next() {
			c <- it.Value()
		}
	}()
	return c
}
