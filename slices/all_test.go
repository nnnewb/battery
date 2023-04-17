package slices

import (
	"github.com/nnnewb/battery/internal/predicate"
	"testing"
)

func TestSlice_All(t *testing.T) {
	type sliceAllArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceAllTestCase[T any] struct {
		name string
		s    Slice[T]
		args sliceAllArgs[T]
		want bool
	}
	tests := []sliceAllTestCase[int]{
		{
			name: "nil slice",
			s:    nil,
			args: sliceAllArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
		{
			name: "all ones",
			s:    []int{1, 1, 1, 1},
			args: sliceAllArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: true,
		},
		{
			name: "all zeros",
			s:    []int{0, 0, 0, 0},
			args: sliceAllArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.All(tt.args.predicate); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
