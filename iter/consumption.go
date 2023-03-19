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

// Reduce 对指定迭代器逐个元素调用函数 f ，函数 f 接收两个参数，一个结果参数 R
// 和迭代元素 T ，返回结果元素 R 。在调用 Reduce 时可指定 R 的初始值。
func Reduce[T, R any](iterator Iterator[T], initial R, f func(R, T) R) R {
	r := initial
	for iterator = iterator.Next(); !iterator.Exhausted(); iterator = iterator.Next() {
		r = f(r, iterator.Value())
	}
	return r
}

// Fold 是 Reduce 的别名
func Fold[T, R any](iterator Iterator[T], initial R, f func(R, T) R) R {
	return Reduce(iterator, initial, f)
}
