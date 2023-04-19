package slices

import (
	"github.com/nnnewb/battery/internal/constraints"
	"sort"
)

func Min[T constraints.Ordered](s Slice[T]) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return Sorted(s)[0], true
}

func Max[T constraints.Ordered](s Slice[T]) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	sorted := Sorted(s)
	return sorted[len(sorted)-1], true
}

func MinKeyFunc[T any, K constraints.Ordered](s Slice[T], keyFunc func(T) K) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	sorted := SortedKeyFunc(s, keyFunc)
	return sorted[0], true
}

func MaxKeyFunc[T any, K constraints.Ordered](s Slice[T], keyFunc func(T) K) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	sorted := SortedKeyFunc(s, keyFunc)
	return sorted[len(sorted)-1], true
}

func (s Slice[T]) MinLessFunc(less func(i, j int) bool) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	r := make(Slice[T], 0, len(s))
	for _, t := range s {
		r = append(r, t)
	}
	sort.Slice(r, less)
	return r[0], true
}

func (s Slice[T]) MaxLessFunc(less func(i, j int) bool) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	r := make(Slice[T], 0, len(s))
	for _, t := range s {
		r = append(r, t)
	}
	sort.Slice(r, less)
	return r[len(r)-1], true
}
