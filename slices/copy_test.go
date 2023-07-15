//go:build go1.18

package slices

import "testing"

// Unit test for the Copy method
func TestCopy(t *testing.T) {
	// create a slice with some elements
	s := []int{1, 2, 3, 4, 5}
	// make a copy of the slice
	copySlice := Copy(s)
	// check that the copy has the same elements as the original slice
	for i := range s {
		if s[i] != copySlice[i] {
			t.Errorf("Copy failed: expected %v, got %v", s, copySlice)
			break
		}
	}
}
