//go:build go1.18

package slices

import "testing"

// We define a test function for the Choices method.
func TestChoices(t *testing.T) {
	// We create a new slice of integers.
	s := []int{1, 2, 3, 4, 5}

	// We call the Choices method with n = 3.
	result := Choices(s, 3)

	// We check that the length of the result slice is 3.
	if len(result) != 3 {
		t.Errorf("Expected length %d, but got %d", 3, len(result))
	}

	// We check that all elements in the result slice are in the original slice.
	for _, element := range result {
		if !Contains(s, element) {
			t.Errorf("Element %d not found in original slice", element)
		}
	}
}
