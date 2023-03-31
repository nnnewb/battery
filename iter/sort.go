package iter

import (
	"github.com/nnnewb/battery/internal/constraints"
	"sort"
)

type sortPair[K constraints.Ordered, V any] struct {
	key K
	val V
}

// Sort 对元素按升序进行排序，排序结果迭代器形式返回。
func Sort[T any, K constraints.Ordered](it Iterator[T], keyFunc func(T) K) Iterator[T] {
	slice := Collect(Map[T, sortPair[K, T]](it, func(t T) sortPair[K, T] {
		return sortPair[K, T]{
			key: keyFunc(it.Value()),
			val: it.Value(),
		}
	}))

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].key < slice[j].key
	})

	return Map(Lift(slice), func(t sortPair[K, T]) T {
		return t.val
	})
}
