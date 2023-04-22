package slices

import (
	"testing"
)

// Unit test for the DropWhile function
func TestSlice_DropWhile(t *testing.T) {
	// Define a test case struct
	type testCase struct {
		name      string
		slice     Slice[int]
		predicate func(int) bool
		expected  Slice[int]
	}

	// Define test cases
	testCases := []testCase{
		{
			name:      "drop while all elements satisfy predicate",
			slice:     Slice[int]{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 6 },
			expected:  Slice[int]{},
		},
		{
			name:      "drop while no elements satisfy predicate",
			slice:     Slice[int]{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 0 },
			expected:  Slice[int]{1, 2, 3, 4, 5},
		},
		{
			name:      "drop while some elements satisfy predicate",
			slice:     Slice[int]{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 3 },
			expected:  Slice[int]{3, 4, 5},
		},
	}

	// Iterate over test cases and run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the DropWhile function with the test case's slice and predicate
			result := tc.slice.DropWhile(tc.predicate)

			// Check that the result matches the expected output
			if !result.Equal(tc.expected, func(i int, i2 int) bool { return i == i2 }) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
