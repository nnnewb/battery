package slices

import "testing"

// Unit test for Shuffle function
func TestShuffle(t *testing.T) {
	// Create a slice of integers
	s := Slice[int]{1, 2, 3, 4, 5}

	// Shuffle the slice
	shuffled := s.Shuffle()

	if shuffled.Equal(s, func(i int, i2 int) bool { return i == i2 }) {
		t.Errorf("Shuffled slice %v equals to original slice %v", shuffled, s)
	}

	// Check that the length of the shuffled slice is the same as the original slice
	if len(shuffled) != len(s) {
		t.Errorf("Shuffled slice length = %v, want %v", len(shuffled), len(s))
	}

	// Check that the shuffled slice contains the same elements as the original slice
	for _, v := range s {
		if !shuffled.Contains(v, func(i int, i2 int) bool { return i == i2 }) {
			t.Errorf("Shuffled slice does not contain element %v", v)
		}
	}
}
