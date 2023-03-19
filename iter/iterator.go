package iter

type Iterator[T any] interface {
	// Next 下一个迭代
	Next() Iterator[T]
	// Value 返回迭代器当前值
	Value() T
	// Exhausted 返回迭代器是否已穷尽
	Exhausted() bool
}
