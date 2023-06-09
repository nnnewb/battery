package iter

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type uniqueArgs[T comparable] struct {
		it Iterator[T]
	}
	type uniqueTestCase[T comparable] struct {
		name string
		args uniqueArgs[T]
		want []T
	}
	tests := []uniqueTestCase[int]{
		{
			name: "exhausted",
			args: uniqueArgs[int]{
				it: Exhausted[int](),
			},
			want: []int{},
		},
		{
			name: "simple",
			args: uniqueArgs[int]{
				it: Lift([]int{1, 1, 3, 3, 5, 5, 1, 1, 3, 3, 5, 5}),
			},
			want: []int{1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.args.it)
			s := Collect(got)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Unique() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestUniqueBy(t *testing.T) {
	type args[T any, K comparable] struct {
		it Iterator[T]
		fn func(T) K
	}
	tests := []struct {
		name string
		args args[int, int]
		want []int
	}{
		{
			name: "exhausted",
			args: args[int, int]{
				it: Exhausted[int](),
				fn: func(i int) int { return i },
			},
			want: []int{},
		},
		{
			name: "simple",
			args: args[int, int]{
				it: Lift([]int{1, 2, 3, 4, 5, 6}),
				fn: func(i int) int { return i % 3 },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueBy(tt.args.it, tt.args.fn)
			s := Collect(got)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("UniqueBy() = %v, want %v", s, tt.want)
			}
		})
	}
}
