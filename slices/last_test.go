package slices

import (
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func TestLast(t *testing.T) {
	s := []int{1, 2, 3, 2, 4, 5, 2}
	// Positive test case
	if idx, ok := Last(s, 2); !ok || idx != 6 {
		t.Errorf("Last(s, 2) returned (%v, %v), expected (6, true)", idx, ok)
	}
	// Negative test case
	if idx, ok := Last(s, 6); ok {
		t.Errorf("Last(s, 6) returned (%v, %v), expected (0, false)", idx, ok)
	}
	// Test case for nil slice
	var nilSlice []int
	if idx, ok := Last(nilSlice, 2); ok {
		t.Errorf("Last(nilSlice, 2) returned (%v, %v), expected (0, false)", idx, ok)
	}
}

func TestLastFunc(t *testing.T) {
	type sliceLastArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceLastTC[T any] struct {
		name string
		s    []T
		args sliceLastArgs[T]
		want T
		ok   bool
	}
	tests := []sliceLastTC[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceLastArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 0,
			ok:   false,
		},
		{
			name: "simple",
			s:    []int{5, 4, 3, 2, 1, 0},
			args: sliceLastArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 4,
			ok:   true,
		},
		{
			name: "simple",
			s:    []int{0, 5, 4, 3, 2, 1, 0},
			args: sliceLastArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: 5,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LastFunc(tt.s, tt.args.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Last() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
