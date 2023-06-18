package slices

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type reduceArgs[T any, R any] struct {
		s       []T
		initial R
		f       func(R, T) R
	}
	type testCase[T any, R any] struct {
		name string
		args reduceArgs[T, R]
		want R
	}
	sum := func(i, i2 int) int { return i + i2 }
	tests := []testCase[int, int]{
		{
			name: "nil",
			args: reduceArgs[int, int]{
				s:       nil,
				initial: 0,
				f:       sum,
			},
			want: 0,
		},
		{
			name: "simple",
			args: reduceArgs[int, int]{
				s:       []int{1, 2, 3},
				initial: 0,
				f:       sum,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.s, tt.args.initial, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
