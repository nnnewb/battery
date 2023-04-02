package slices

import (
	"github.com/nnnewb/battery/internal/constraints"
	"github.com/nnnewb/battery/iter"
	"sort"
)

type sortPair[K constraints.Ordered, V any] struct {
	key K
	val V
}

func Sorted[T constraints.Ordered](s Slice[T]) Slice[T] {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func SortedKeyFunc[T any, K constraints.Ordered](s Slice[T], keyFunc func(T) K) Slice[T] {
	slice := iter.Collect(iter.Map(iter.Lift(s), func(t T) sortPair[K, T] {
		return sortPair[K, T]{
			key: keyFunc(t),
			val: t,
		}
	}))

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].key < slice[j].key
	})

	return iter.Collect(iter.Map(iter.Lift(slice), func(t sortPair[K, T]) T {
		return t.val
	}))
}

func SortedLessFunc[T any](s Slice[T], less func(i, j int) bool) Slice[T] {
	sort.Slice(s, less)
	return s
}
