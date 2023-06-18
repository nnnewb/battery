package slices

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	type groupByArgs[T any, K comparable] struct {
		s       []T
		keyFunc func(T) K
	}
	type groupByTC[T any, K comparable] struct {
		name string
		args groupByArgs[T, K]
		want map[K][]T
	}
	tests := []groupByTC[int, int]{
		{
			name: "nil",
			args: groupByArgs[int, int]{
				s: nil,
				keyFunc: func(i int) int {
					return i
				},
			},
			want: map[int][]int{},
		},
		{
			name: "simple",
			args: groupByArgs[int, int]{
				s: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 4, 3, 2, 1},
				keyFunc: func(i int) int {
					return i
				},
			},
			want: map[int][]int{
				1: {1, 1, 1},
				2: {2, 2, 2},
				3: {3, 3, 3},
				4: {4, 4, 4},
				5: {5, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.s, tt.args.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
