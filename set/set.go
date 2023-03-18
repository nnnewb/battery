package set

import (
	"github.com/nnnewb/battery/iter"
)

type Set[T comparable] map[T]bool

// Add 往集合添加一个元素。
//
//	s := make(Set[int])
//	s.Add(1,2,3)
func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = true
	}
}

// Remove 从集合删除一个元素
//
//	s:=make(Set[int])
//	s.Add(1,2,3)
//	s.Remove(3)
//	println(s.Contains(3))
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Contains 返回指定元素是否存在于集合中
//
//	s := make(Set[int])
//	s.Add(1,2,3)
//	s.Contains(1)
func (s Set[T]) Contains(v T) bool {
	return s[v]
}

// RelativeComplement 对集合做差运算，返回属于本集合但不属于 other
// 集合的元素集合。
//
//	s1 := make(Set[int])
//	s2 := make(Set[int])
//	s1.Add(1,2,3)
//	s2.Add(2,3,4)
//	r := s1.RelativeComplement(s2)
//	r.Contains(1)
func (s Set[T]) RelativeComplement(other Set[T]) Set[T] {
	ret := make(Set[T])
	for v := range s {
		if !other.Contains(v) {
			ret.Add(v)
		}
	}
	return ret
}

// Union 取两个集合的合集。返回两个集合元素的集合。
//
//	s1 := make(Set[int])
//	s2 := make(Set[int])
//	s1.Add(1,2)
//	s2.Add(3,4)
//	r := s1.Union(s2)
//	r.Contains(1)
//	r.Contains(3)
func (s Set[T]) Union(other Set[T]) Set[T] {
	ret := make(Set[T])
	for v := range s {
		ret.Add(v)
	}
	for v := range other {
		ret.Add(v)
	}
	return ret
}

// Intersection 取两个集合的交集。返回两个集合都存在的元素。
//
//	s1 := make(Set[int])
//	s2 := make(Set[int])
//	s1.Add(1,2)
//	s2.Add(2,3)
//	r := s1.Intersection(s2)
//	r.Contains(2)
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	ret := make(Set[T])
	for v := range s {
		if other.Contains(v) {
			ret.Add(v)
		}
	}
	return ret
}

// Copy 复制集合。
//
//	s := make(Set[int])
//	s.Add(1,2)
//	s2 := s.Copy()
//	s2.Contains(1)
func (s Set[T]) Copy() Set[T] {
	ret := make(Set[T])
	for v := range s {
		ret.Add(v)
	}
	return ret
}

// ToSlice 把集合转为切片返回。
//
//	s := make(Set[int])
//	s.Add(1,2)
//	slice := s.ToSlice()
//	println(len(slice) == 2)
func (s Set[T]) ToSlice() []T {
	return iter.Collect(s.Iter())
}

// Iter 基于 Chan 创建的迭代器。
//
//	s := make(Set[int])
//	s.Add(1,2)
//	slice := iter.Collect(iter.Filter(s.Iter(), func(i int) bool { return i > 1}))
//	println(len(slice)==1)
func (s Set[T]) Iter() iter.Iterator[T] {
	return iter.FromGenerateFunc(func(exhausted func()) T {
		var t T
		for v := range s {
			return v
		}
		exhausted()
		return t
	})
}
