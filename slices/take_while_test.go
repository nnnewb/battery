package slices

import "testing"

// TestTakeWhile tests the TakeWhile function
func TestTakeWhile(t *testing.T) {
	// Define a test case struct
	tests := []struct {
		name      string
		slice     Slice[int]
		predicate func(int) bool
		want      Slice[int]
	}{
		{
			name:      "empty slice",
			slice:     Slice[int]{},
			predicate: func(i int) bool { return i < 5 },
			want:      Slice[int]{},
		},
		{
			name:      "all elements satisfy predicate",
			slice:     Slice[int]{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 6 },
			want:      Slice[int]{1, 2, 3, 4, 5},
		},
		{
			name:      "no elements satisfy predicate",
			slice:     Slice[int]{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i > 5 },
			want:      Slice[int]{},
		},
		{
			name:      "some elements satisfy predicate",
			slice:     Slice[int]{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 3 },
			want:      Slice[int]{1, 2},
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		// Call TakeWhile on the test slice with the test predicate
		got := tt.slice.TakeWhile(tt.predicate)

		// Check that the returned slice matches the expected slice
		if !got.Equal(tt.want, func(i int, i2 int) bool { return i == i2 }) {
			t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
		}
	}
}
