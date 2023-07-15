//go:build go1.18

package slices

import (
	"reflect"
	"testing"

	"github.com/nnnewb/battery/internal/predicate"
)

func TestFilter(t *testing.T) {
	type sliceFilterArgs[T any] struct {
		predicate func(T) bool
	}
	type sliceFilterTestCase[T any] struct {
		name string
		s    []T
		args sliceFilterArgs[T]
		want []T
	}
	tests := []sliceFilterTestCase[int]{
		{
			name: "nil slice parameter",
			s:    nil,
			args: sliceFilterArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: []int{},
		},
		{
			name: "non-zero result",
			s:    []int{1, 2, 3, 4, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5},
			args: sliceFilterArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
		},
		{
			name: "no result",
			s:    []int{-1, -2, -3, -4, -5, -4, -3, -2, -1, 0, -1, -2, -3, -4, -5},
			args: sliceFilterArgs[int]{
				predicate: predicate.IsPositive[int],
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.s, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("case \"%s\": Filter() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
