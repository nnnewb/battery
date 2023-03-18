package iter

type Iterator[T any] interface {
	Next() Iterator[T]
	Value() T
	Exhausted() bool
}
