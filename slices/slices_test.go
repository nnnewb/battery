package slices

import "testing"

func TestSliceTypeConversion(t *testing.T) {
	ordinary := Slice[int]{1, 2, 3}
	_ = OrderedElemSlice[int](ordinary)
	_ = ComparableElemSlice[int](ordinary)
}
