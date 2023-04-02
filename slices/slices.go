package slices

import "github.com/nnnewb/battery/internal/constraints"

// Slice 任意切片类型
type Slice[T any] []T

// ComparableElemSlice 可比较元素的切片
type ComparableElemSlice[T comparable] []T

// OrderedElemSlice 可排序元素的切片
type OrderedElemSlice[T constraints.Ordered] []T

// TODO 期望把 golang.org/x/exp/slices 下的内容实现一遍
