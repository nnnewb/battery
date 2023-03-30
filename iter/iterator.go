package iter

type Iterator[T any] interface {
	// Next 下一个迭代
	Next() bool
	// Value 返回迭代器当前值
	Value() T
}
