package slices

import (
	"reflect"
	"testing"
)

func TestCountBy(t *testing.T) {
	type countByArgs[T any, K comparable] struct {
		s       []T
		keyFunc func(T) K
	}
	type countByTestCase[T any, K comparable] struct {
		name string
		args countByArgs[T, K]
		want map[K]int
	}
	tests := []countByTestCase[int, int]{
		{
			name: "nil",
			args: countByArgs[int, int]{
				s: nil,
				keyFunc: func(i int) int {
					return i
				},
			},
			want: map[int]int{},
		},
		{
			name: "simple",
			args: countByArgs[int, int]{
				s: []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3},
				keyFunc: func(i int) int {
					return i
				},
			},
			want: map[int]int{
				1: 3,
				2: 3,
				3: 3,
				4: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBy(tt.args.s, tt.args.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
