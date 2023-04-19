package slices

import (
	"sort"
)

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

func (s ComparableElemSlice[T]) MinLessFunc(less func(i, j int) bool) (T, bool) {
	return Slice[T](s).MinLessFunc(less)
}

func (s ComparableElemSlice[T]) MaxLessFunc(less func(i, j int) bool) (T, bool) {
	return Slice[T](s).MaxLessFunc(less)
}

func (s OrderedElemSlice[T]) Max() (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	r := make(OrderedElemSlice[T], len(s))
	copy(r, s)
	sort.Sort(r)
	return r[len(r)-1], true
}

func (s OrderedElemSlice[T]) Min() (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	r := make(OrderedElemSlice[T], len(s))
	copy(r, s)
	sort.Sort(r)
	return r[0], true
}
