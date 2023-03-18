package iter

import "github.com/nnnewb/battery/optional"

// FromSlice 创建一个切片迭代器
func FromSlice[T any](s []T) Iterator[T] {
	return sliceIter[T]{
		underlying: s,
	}
}

// FromChan 从 go 内置的 chan 创建 Iterator
func FromChan[T any](c <-chan optional.Optional[T]) Iterator[T] {
	i := chanIter[T]{c: c}
	return i.Next()
}

// FromGenerateFunc 从生成函数创建迭代器
func FromGenerateFunc[T any](f func(exhausted func()) T) Iterator[T] {
	return FromChan(gen(f))
}

type KeyValue[T comparable, U any] struct {
	Key   T
	Value U
}

func FromMap[T comparable, U any](m map[T]U) Iterator[KeyValue[T, U]] {
	return FromChan(gen(func(exhausted func()) KeyValue[T, U] {
		for k, v := range m {
			return KeyValue[T, U]{
				Key:   k,
				Value: v,
			}
		}
		exhausted()
		return KeyValue[T, U]{}
	}))
}
