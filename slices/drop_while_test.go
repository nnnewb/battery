package slices

import (
	"testing"
)

// Unit test for the DropWhile function
func TestDropWhile(t *testing.T) {
	// Define a test case struct
	type testCase struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  []int
	}

	// Define test cases
	testCases := []testCase{
		{
			name:      "drop while all elements satisfy predicate",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 6 },
			expected:  []int{},
		},
		{
			name:      "drop while no elements satisfy predicate",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 0 },
			expected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "drop while some elements satisfy predicate",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 3 },
			expected:  []int{3, 4, 5},
		},
	}

	// Iterate over test cases and run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the DropWhile function with the test case's slice and predicate
			result := DropWhile(tc.slice, tc.predicate)

			// Check that the result matches the expected output
			if !Equal(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
