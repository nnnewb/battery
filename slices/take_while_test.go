package slices

import "testing"

// TestTakeWhile tests the TakeWhile function
func TestTakeWhile(t *testing.T) {
	// Define a test case struct
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(i int) bool { return i < 5 },
			want:      []int{},
		},
		{
			name:      "all elements satisfy predicate",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 6 },
			want:      []int{1, 2, 3, 4, 5},
		},
		{
			name:      "no elements satisfy predicate",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i > 5 },
			want:      []int{},
		},
		{
			name:      "some elements satisfy predicate",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 3 },
			want:      []int{1, 2},
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		// Call TakeWhile on the test slice with the test predicate
		got := TakeWhile(tt.slice, tt.predicate)

		// Check that the returned slice matches the expected slice
		if !Equal(got, tt.want) {
			t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
		}
	}
}
