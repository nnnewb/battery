package slices

import "testing"

// TestReverse tests the Reverse function
func TestReverse(t *testing.T) {
	// Test with an empty slice
	s := make([]int, 0)
	reversed := Reverse(s)
	if len(reversed) != 0 {
		t.Errorf("Reverse of an empty slice should be an empty slice")
	}

	// Test with a slice of length 1
	s = []int{1}
	reversed = Reverse(s)
	if len(reversed) != 1 || reversed[0] != 1 {
		t.Errorf("Reverse of a slice of length 1 should be the same slice")
	}

	// Test with a slice of length 2
	s = []int{1, 2}
	reversed = Reverse(s)
	if len(reversed) != 2 || reversed[0] != 2 || reversed[1] != 1 {
		t.Errorf("Reverse of a slice of length 2 should be [2, 1]")
	}

	// Test with a slice of length 3
	s = []int{1, 2, 3}
	reversed = Reverse(s)
	if len(reversed) != 3 || reversed[0] != 3 || reversed[1] != 2 || reversed[2] != 1 {
		t.Errorf("Reverse of a slice of length 3 should be [3, 2, 1]")
	}
}
