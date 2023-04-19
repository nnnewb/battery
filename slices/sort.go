package slices

import (
	"sort"
)

func (s Slice[T]) SortLessFunc(less func(i, j int) bool) Slice[T] {
	sort.Slice(s, less)
	return s
}

func (s ComparableElemSlice[T]) SortLessFunc(less func(i, j int) bool) ComparableElemSlice[T] {
	return ComparableElemSlice[T](Slice[T](s).SortLessFunc(less))
}

func (s OrderedElemSlice[T]) Sort() OrderedElemSlice[T] {
	sort.Sort(s)
	return s
}
