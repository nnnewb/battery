package slices

import "github.com/nnnewb/battery/internal/constraints"

// Slice 任意切片类型
type Slice[T any] []T

// ComparableElemSlice 可比较元素的切片
type ComparableElemSlice[T comparable] []T

// OrderedElemSlice 可排序元素的切片
type OrderedElemSlice[T constraints.Ordered] []T

func (s OrderedElemSlice[T]) Len() int {
	return len(s)
}

func (s OrderedElemSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s OrderedElemSlice[T]) Swap(i, j int) {
	var t T
	t = s[i]
	s[i] = s[j]
	s[j] = t
}

// TODO 期望把 golang.org/x/exp/slices 下的内容实现一遍
