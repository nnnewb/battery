//go:build go1.18

package slices

import "testing"

// Unit test for Shuffle function
func TestShuffle(t *testing.T) {
	// Create a slice of integers
	s := []int{1, 2, 3, 4, 5}

	// Shuffle the slice
	shuffled := Shuffle(s)

	if Equal(shuffled, s) {
		t.Errorf("Shuffled slice %v equals to original slice %v", shuffled, s)
	}

	// Check that the length of the shuffled slice is the same as the original slice
	if len(shuffled) != len(s) {
		t.Errorf("Shuffled slice length = %v, want %v", len(shuffled), len(s))
	}

	// Check that the shuffled slice contains the same elements as the original slice
	for _, v := range s {
		if !Contains(shuffled, v) {
			t.Errorf("Shuffled slice does not contain element %v", v)
		}
	}
}
