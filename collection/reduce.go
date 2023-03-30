package collection

// Reduce 对指定迭代器逐个元素调用函数 f ，函数 f 接收两个参数，一个结果参数 R
// 和迭代元素 T ，返回结果元素 R 。在调用 Reduce 时可指定 R 的初始值。
func Reduce[T, R any](iterable Iterable[T], initial R, f func(R, T) R) R {
	r := initial
	it := iterable.Iterator()
	for it.Next() {
		r = f(r, it.Value())
	}
	return r
}

// Fold 是 Reduce 的别名
func Fold[T, R any](iterable Iterable[T], initial R, f func(R, T) R) R {
	return Reduce(iterable, initial, f)
}
