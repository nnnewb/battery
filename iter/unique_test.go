package iter

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type uniqueArgs[T comparable] struct {
		it Iterator[int]
	}
	type uniqueTestCase[T comparable] struct {
		name string
		args uniqueArgs[T]
		want []int
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
