package slices

import (
	"github.com/nnnewb/battery/internal/predicate"
	"reflect"
	"testing"
)

func TestSlice_Filter(t *testing.T) {
	type sliceFilterArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceFilterTestCase[T any] struct {
		name string
		s    Slice[T]
		args sliceFilterArgs[T]
		want Slice[T]
	}
	tests := []sliceFilterTestCase[int]{
		{
			name: "nil",
			s:    nil,
			args: sliceFilterArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: []int{},
		},
		{
			name: "simple",
			s:    []int{1, 2, 3, 4, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5},
			args: sliceFilterArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Filter(tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
