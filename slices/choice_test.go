package slices

import "testing"

// We define a test function for the Choices method.
func TestChoices(t *testing.T) {
	// We create a new slice of integers.
	s := Slice[int]{1, 2, 3, 4, 5}

	// We call the Choices method with n = 3.
	result := s.Choices(3)

	// We check that the length of the result slice is 3.
	if len(result) != 3 {
		t.Errorf("Expected length %d, but got %d", 3, len(result))
	}

	// We check that all elements in the result slice are in the original slice.
	for _, element := range result {
		if !s.Contains(element, func(i int, i2 int) bool { return i == i2 }) {
			t.Errorf("Element %d not found in original slice", element)
		}
	}
}
