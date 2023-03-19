package iter

import "math"

type countIter struct {
	i int
}

func (it countIter) Next() Iterator[int] {
	if it.i < math.MaxInt {
		it.i++
	}
	return it
}

func (it countIter) Value() int {
	return it.i
}

func (it countIter) Exhausted() bool {
	return math.MaxInt == it.i
}

// Count 返回迭代器，返回从 1 开始递增的正整数。最大不超过 math.MaxInt 。
func Count() Iterator[int] {
	return countIter{
		i: 0,
	}
}
