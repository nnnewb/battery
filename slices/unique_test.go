package slices

import (
	"reflect"
	"testing"
)

func TestComparableElemSlice_Unique(t *testing.T) {
	testCases := []struct {
		name     string
		input    ComparableElemSlice[int]
		expected []int
	}{
		{
			name:     "no duplicates",
			input:    ComparableElemSlice[int]{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "some duplicates",
			input:    ComparableElemSlice[int]{1, 2, 2, 3, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all duplicates",
			input:    ComparableElemSlice[int]{1, 1, 1, 1},
			expected: []int{1},
		},
		{
			name:     "empty slice",
			input:    ComparableElemSlice[int]{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := []int(tc.input.Unique())
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
