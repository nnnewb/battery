package slices

import (
	"github.com/nnnewb/battery/internal/predicate"
	"testing"
)

func TestSlice_Any(t *testing.T) {
	type sliceAnyArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceAnyTestCase[T any] struct {
		name string
		s    Slice[T]
		args sliceAnyArgs[T]
		want bool
	}
	tests := []sliceAnyTestCase[int]{
		{
			name: "nil slice",
			s:    nil,
			args: sliceAnyArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
		{
			name: "all ones",
			s:    []int{1, 1, 1, 1, 1},
			args: sliceAnyArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: true,
		},
		{
			name: "all zeros",
			s:    []int{0, 0, 0, 0, 0},
			args: sliceAnyArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: false,
		},
		{
			name: "mixed ones zeros",
			s:    []int{1, 0, 1, 0, 1, 0},
			args: sliceAnyArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Any(tt.args.predicate); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}
